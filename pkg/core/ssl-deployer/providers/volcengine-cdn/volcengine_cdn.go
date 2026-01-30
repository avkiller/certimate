package volcenginecdn

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	vecdn "github.com/volcengine/volc-sdk-golang/service/cdn"
	ve "github.com/volcengine/volcengine-go-sdk/volcengine"

	"github.com/certimate-go/certimate/pkg/core"
	sslmgrsp "github.com/certimate-go/certimate/pkg/core/ssl-manager/providers/volcengine-cdn"
	xcert "github.com/certimate-go/certimate/pkg/utils/cert"
)

type SSLDeployerProviderConfig struct {
	// 火山引擎 AccessKeyId。
	AccessKeyId string `json:"accessKeyId"`
	// 火山引擎 AccessKeySecret。
	AccessKeySecret string `json:"accessKeySecret"`
	// 域名匹配模式。
	// 零值时默认值 [MatchPatternExact]。
	MatchPattern string `json:"matchPattern,omitempty"`
	// 加速域名（支持泛域名）。
	Domain string `json:"domain"`
}

type SSLDeployerProvider struct {
	config     *SSLDeployerProviderConfig
	logger     *slog.Logger
	sdkClient  *vecdn.CDN
	sslManager core.SSLManager
}

var _ core.SSLDeployer = (*SSLDeployerProvider)(nil)

func NewSSLDeployerProvider(config *SSLDeployerProviderConfig) (*SSLDeployerProvider, error) {
	if config == nil {
		return nil, errors.New("the configuration of the ssl deployer provider is nil")
	}

	client := vecdn.NewInstance()
	client.Client.SetAccessKey(config.AccessKeyId)
	client.Client.SetSecretKey(config.AccessKeySecret)

	sslmgr, err := sslmgrsp.NewSSLManagerProvider(&sslmgrsp.SSLManagerProviderConfig{
		AccessKeyId:     config.AccessKeyId,
		AccessKeySecret: config.AccessKeySecret,
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

	d.sslManager.SetLogger(logger)
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

	// 获取待部署的 CDN 实例
	domains := make([]string, 0)
	switch d.config.MatchPattern {
	case "", MatchPatternExact:
		{
			domains = append(domains, d.config.Domain)
		}

	case MatchPatternWildcard:
		{
			if strings.HasPrefix(d.config.Domain, "*.") {
				temp, err := d.getMatchedDomainsByWildcard(ctx, d.config.Domain)
				if err != nil {
					return nil, err
				}

				domains = temp
			} else {
				domains = append(domains, d.config.Domain)
			}
		}

	case MatchPatternCertSAN:
		{
			temp, err := d.getMatchedDomainsByCertId(ctx, upres.CertId)
			if err != nil {
				return nil, err
			}

			domains = temp
		}

	default:
		return nil, fmt.Errorf("unsupported match pattern: '%s'", d.config.MatchPattern)
	}

	// 遍历绑定证书
	if len(domains) == 0 {
		d.logger.Info("no cdn domains to deploy")
	} else {
		d.logger.Info("found cdn domains to deploy", slog.Any("domains", domains))
		var errs []error

		for _, domain := range domains {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
				if err := d.bindCert(ctx, domain, upres.CertId); err != nil {
					errs = append(errs, err)
				}
			}
		}

		if len(errs) > 0 {
			return nil, errors.Join(errs...)
		}
	}

	return &core.SSLDeployResult{}, nil
}

func (d *SSLDeployerProvider) getMatchedDomainsByWildcard(ctx context.Context, wildcardDomain string) ([]string, error) {
	domains := make([]string, 0)

	// 遍历获取加速域名列表，获取匹配的域名
	// REF: https://www.volcengine.com/docs/6454/75269
	listCdnDomainsPageNum := int64(1)
	listCdnDomainsPageSize := int64(100)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		listCdnDomainsReq := &vecdn.ListCdnDomainsRequest{
			Domain:   ve.String(strings.TrimPrefix(wildcardDomain, "*.")),
			Status:   ve.String("online"),
			PageNum:  ve.Int64(listCdnDomainsPageNum),
			PageSize: ve.Int64(listCdnDomainsPageSize),
		}
		listCdnDomainsResp, err := d.sdkClient.ListCdnDomains(listCdnDomainsReq)
		d.logger.Debug("sdk request 'cdn.ListCdnDomains'", slog.Any("request", listCdnDomainsReq), slog.Any("response", listCdnDomainsResp))
		if err != nil {
			return nil, fmt.Errorf("failed to execute sdk request 'cdn.ListCdnDomains': %w", err)
		}

		if listCdnDomainsResp.Result.Data != nil {
			for _, domain := range listCdnDomainsResp.Result.Data {
				if xcert.MatchHostname(wildcardDomain, domain.Domain) {
					domains = append(domains, domain.Domain)
				}
			}
		}

		if len(listCdnDomainsResp.Result.Data) < int(listCdnDomainsPageSize) {
			break
		} else {
			listCdnDomainsPageSize++
		}
	}

	return domains, nil
}

func (d *SSLDeployerProvider) getMatchedDomainsByCertId(ctx context.Context, cloudCertId string) ([]string, error) {
	domains := make([]string, 0)

	// 获取指定证书可关联的域名
	// REF: https://www.volcengine.com/docs/6454/125711
	describeCertConfigReq := &vecdn.DescribeCertConfigRequest{
		CertId: cloudCertId,
	}
	describeCertConfigResp, err := d.sdkClient.DescribeCertConfig(describeCertConfigReq)
	d.logger.Debug("sdk request 'cdn.DescribeCertConfig'", slog.Any("request", describeCertConfigReq), slog.Any("response", describeCertConfigResp))
	if err != nil {
		return nil, fmt.Errorf("failed to execute sdk request 'cdn.DescribeCertConfig': %w", err)
	}

	if describeCertConfigResp.Result.CertNotConfig != nil {
		for i := range describeCertConfigResp.Result.CertNotConfig {
			domains = append(domains, describeCertConfigResp.Result.CertNotConfig[i].Domain)
		}
	}

	if describeCertConfigResp.Result.OtherCertConfig != nil {
		for i := range describeCertConfigResp.Result.OtherCertConfig {
			domains = append(domains, describeCertConfigResp.Result.OtherCertConfig[i].Domain)
		}
	}

	if len(domains) == 0 {
		if len(describeCertConfigResp.Result.SpecifiedCertConfig) == 0 {
			return nil, errors.New("domains not found")
		}
	}

	return domains, nil
}

func (d *SSLDeployerProvider) bindCert(ctx context.Context, domain string, cloudCertId string) error {
	// 关联证书与加速域名
	// REF: https://www.volcengine.com/docs/6454/125712
	batchDeployCertReq := &vecdn.BatchDeployCertRequest{
		CertId: cloudCertId,
		Domain: domain,
	}
	batchDeployCertResp, err := d.sdkClient.BatchDeployCert(batchDeployCertReq)
	d.logger.Debug("sdk request 'cdn.BatchDeployCert'", slog.Any("request", batchDeployCertReq), slog.Any("response", batchDeployCertResp))
	if err != nil {
		return err
	}

	return nil
}
