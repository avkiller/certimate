package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	edgioapps "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/edgio-applications"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeEdgioApplications, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForEdgio{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := edgioapps.NewSSLDeployerProvider(&edgioapps.SSLDeployerProviderConfig{
			ClientId:      credentials.ClientId,
			ClientSecret:  credentials.ClientSecret,
			EnvironmentId: xmaps.GetString(options.ProviderExtendedConfig, "environmentId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
