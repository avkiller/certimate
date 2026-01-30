package aliyuncdn

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	alicdn "github.com/alibabacloud-go/cdn-20180510/v8/client"
	aliopen "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/samber/lo"

	"github.com/certimate-go/certimate/pkg/core"
	sslmgrsp "github.com/certimate-go/certimate/pkg/core/ssl-manager/providers/aliyun-cas"
)

type SSLDeployerProviderConfig struct {
	// 阿里云 AccessKeyId。
	AccessKeyId string `json:"accessKeyId"`
	// 阿里云 AccessKeySecret。
	AccessKeySecret string `json:"accessKeySecret"`
	// 阿里云资源组 ID。
	ResourceGroupId string `json:"resourceGroupId,omitempty"`
	// 阿里云地域。
	Region string `json:"region"`
	// 加速域名（支持泛域名）。
	Domain string `json:"domain"`
}

type SSLDeployerProvider struct {
	config     *SSLDeployerProviderConfig
	logger     *slog.Logger
	sdkClient  *alicdn.Client
	sslManager core.SSLManager
}

var _ core.SSLDeployer = (*SSLDeployerProvider)(nil)

func NewSSLDeployerProvider(config *SSLDeployerProviderConfig) (*SSLDeployerProvider, error) {
	if config == nil {
		return nil, errors.New("the configuration of the ssl deployer provider is nil")
	}

	client, err := createSDKClient(config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("could not create sdk client: %w", err)
	}

	sslmgr, err := sslmgrsp.NewSSLManagerProvider(&sslmgrsp.SSLManagerProviderConfig{
		AccessKeyId:     config.AccessKeyId,
		AccessKeySecret: config.AccessKeySecret,
		ResourceGroupId: config.ResourceGroupId,
		Region: lo.
			If(config.Region == "" || strings.HasPrefix(config.Region, "cn-"), "cn-hangzhou").
			Else("ap-southeast-1"),
	})
	if err != nil {
		return nil, fmt.Errorf("could not create ssl manager: %w", err)
	}

	return &SSLDeployerProvider{
		config:     config,
		logger:     slog.Default(),
		sdkClient:  client,
		sslManager: sslmgr,
	}, nil
}

func (d *SSLDeployerProvider) SetLogger(logger *slog.Logger) {
	if logger == nil {
		d.logger = slog.New(slog.DiscardHandler)
	} else {
		d.logger = logger
	}
}

func (d *SSLDeployerProvider) Deploy(ctx context.Context, certPEM string, privkeyPEM string) (*core.SSLDeployResult, error) {
	if d.config.Domain == "" {
		return nil, errors.New("config `domain` is required")
	}

	// 上传证书
	upres, err := d.sslManager.Upload(ctx, certPEM, privkeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to upload certificate file: %w", err)
	} else {
		d.logger.Info("ssl certificate uploaded", slog.Any("result", upres))
	}

	// "*.example.com" → ".example.com"，适配阿里云 CDN 要求的泛域名格式
	domain := strings.TrimPrefix(d.config.Domain, "*")

	// 设置 CDN 域名域名证书
	// REF: https://help.aliyun.com/zh/cdn/developer-reference/api-cdn-2018-05-10-setcdndomainsslcertificate
	certId, _ := strconv.ParseInt(upres.CertId, 10, 64)
	setCdnDomainSSLCertificateReq := &alicdn.SetCdnDomainSSLCertificateRequest{
		DomainName: tea.String(domain),
		CertType:   tea.String("cas"),
		CertId:     tea.Int64(int64(certId)),
		CertRegion: lo.
			If(d.config.Region == "" || strings.HasPrefix(d.config.Region, "cn-"), tea.String("cn-hangzhou")).
			Else(tea.String("ap-southeast-1")),
		SSLProtocol: tea.String("on"),
	}
	setCdnDomainSSLCertificateResp, err := d.sdkClient.SetCdnDomainSSLCertificate(setCdnDomainSSLCertificateReq)
	d.logger.Debug("sdk request 'cdn.SetCdnDomainSSLCertificate'", slog.Any("request", setCdnDomainSSLCertificateReq), slog.Any("response", setCdnDomainSSLCertificateResp))
	if err != nil {
		return nil, fmt.Errorf("failed to execute sdk request 'cdn.SetCdnDomainSSLCertificate': %w", err)
	}

	return &core.SSLDeployResult{}, nil
}

func createSDKClient(accessKeyId, accessKeySecret string) (*alicdn.Client, error) {
	config := &aliopen.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String("cdn.aliyuncs.com"),
	}

	client, err := alicdn.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
