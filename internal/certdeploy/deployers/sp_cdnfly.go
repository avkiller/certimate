package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/cdnfly"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeCdnfly, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForCdnfly{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		deployer, err := cdnfly.NewSSLDeployerProvider(&cdnfly.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiKey:                   credentials.ApiKey,
			ApiSecret:                credentials.ApiSecret,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             cdnfly.ResourceType(xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "resourceType", string(cdnfly.RESOURCE_TYPE_SITE))),
			SiteId:                   xmaps.GetString(options.ProviderExtendedConfig, "siteId"),
			CertificateId:            xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
		})
		return deployer, err
	}); err != nil {
		panic(err)
	}
}
