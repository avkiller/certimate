package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	unicloudwebhost "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/unicloud-webhost"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeUniCloudWebHost, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForUniCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := unicloudwebhost.NewSSLDeployerProvider(&unicloudwebhost.SSLDeployerProviderConfig{
			Username:      credentials.Username,
			Password:      credentials.Password,
			SpaceProvider: xmaps.GetString(options.ProviderExtendedConfig, "spaceProvider"),
			SpaceId:       xmaps.GetString(options.ProviderExtendedConfig, "spaceId"),
			Domain:        xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
