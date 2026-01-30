package volcenginelive

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	velive "github.com/volcengine/volc-sdk-golang/service/live/v20230101"
	ve "github.com/volcengine/volcengine-go-sdk/volcengine"

	"github.com/certimate-go/certimate/pkg/core"
	sslmgrsp "github.com/certimate-go/certimate/pkg/core/ssl-manager/providers/volcengine-live"
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
	// 直播流域名（支持泛域名）。
	Domain string `json:"domain"`
}

type SSLDeployerProvider struct {
	config     *SSLDeployerProviderConfig
	logger     *slog.Logger
	sdkClient  *velive.Live
	sslManager core.SSLManager
}

var _ core.SSLDeployer = (*SSLDeployerProvider)(nil)

func NewSSLDeployerProvider(config *SSLDeployerProviderConfig) (*SSLDeployerProvider, error) {
	if config == nil {
		return nil, errors.New("the configuration of the ssl deployer provider is nil")
	}

	client := velive.NewInstance()
	client.SetAccessKey(config.AccessKeyId)
	client.SetSecretKey(config.AccessKeySecret)

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

	// 获取待部署的直播实例
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

	default:
		return nil, fmt.Errorf("unsupported match pattern: '%s'", d.config.MatchPattern)
	}

	// 遍历绑定证书
	if len(domains) == 0 {
		d.logger.Info("no live domains to deploy")
	} else {
		d.logger.Info("found live domains to deploy", slog.Any("domains", domains))
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

	// 遍历查询域名列表，获取匹配的域名
	// REF: https://www.volcengine.com/docs/6469/1186277#%E6%9F%A5%E8%AF%A2%E5%9F%9F%E5%90%8D%E5%88%97%E8%A1%A8
	listDomainDetailPageNum := int32(1)
	listDomainDetailPageSize := int32(1000)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		listDomainDetailReq := &velive.ListDomainDetailBody{
			DomainStatusList: ve.Int32Slice([]int32{0}),
			PageNum:          listDomainDetailPageNum,
			PageSize:         listDomainDetailPageSize,
		}
		listDomainDetailResp, err := d.sdkClient.ListDomainDetail(ctx, listDomainDetailReq)
		d.logger.Debug("sdk request 'live.ListDomainDetail'", slog.Any("request", listDomainDetailReq), slog.Any("response", listDomainDetailResp))
		if err != nil {
			return nil, fmt.Errorf("failed to execute sdk request 'live.ListDomainDetail': %w", err)
		}

		if listDomainDetailResp.Result.DomainList != nil {
			for _, domain := range listDomainDetailResp.Result.DomainList {
				if xcert.MatchHostname(wildcardDomain, domain.Domain) {
					domains = append(domains, domain.Domain)
				}
			}
		}

		if len(listDomainDetailResp.Result.DomainList) < int(listDomainDetailPageSize) {
			break
		} else {
			listDomainDetailPageNum++
		}
	}

	if len(domains) == 0 {
		return nil, errors.New("domain not found")
	}

	return domains, nil
}

func (d *SSLDeployerProvider) bindCert(ctx context.Context, domain string, cloudCertId string) error {
	// 绑定证书
	// REF: https://www.volcengine.com/docs/6469/1186278#%E7%BB%91%E5%AE%9A%E8%AF%81%E4%B9%A6
	bindCertReq := &velive.BindCertBody{
		ChainID: cloudCertId,
		Domain:  domain,
		HTTPS:   ve.Bool(true),
	}
	bindCertResp, err := d.sdkClient.BindCert(ctx, bindCertReq)
	d.logger.Debug("sdk request 'live.BindCert'", slog.Any("request", bindCertReq), slog.Any("response", bindCertResp))
	if err != nil {
		return err
	}

	return nil
}
