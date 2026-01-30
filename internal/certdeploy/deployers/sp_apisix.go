package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/apisix"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAPISIX, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAPISIX{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := apisix.NewSSLDeployerProvider(&apisix.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiKey:                   credentials.ApiKey,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             apisix.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			CertificateId:            xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
