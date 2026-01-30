package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	wangsucdnpro "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/wangsu-cdnpro"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeWangsuCDNPro, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForWangsu{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := wangsucdnpro.NewSSLDeployerProvider(&wangsucdnpro.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			ApiKey:          credentials.ApiKey,
			Environment:     xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "environment", "production"),
			Domain:          xmaps.GetString(options.ProviderExtendedConfig, "domain"),
			CertificateId:   xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
			WebhookId:       xmaps.GetString(options.ProviderExtendedConfig, "webhookId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
