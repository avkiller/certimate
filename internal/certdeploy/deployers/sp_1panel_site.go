package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	opsite "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/1panel-site"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderType1PanelSite, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigFor1Panel{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := opsite.NewSSLDeployerProvider(&opsite.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiVersion:               credentials.ApiVersion,
			ApiKey:                   credentials.ApiKey,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			NodeName:                 xmaps.GetString(options.ProviderExtendedConfig, "nodeName"),
			ResourceType:             opsite.ResourceType(xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "resourceType", string(opsite.RESOURCE_TYPE_WEBSITE))),
			WebsiteId:                xmaps.GetInt64(options.ProviderExtendedConfig, "websiteId"),
			CertificateId:            xmaps.GetInt64(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
