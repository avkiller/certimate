package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/deployer"
	baotapanelgoconsole "github.com/certimate-go/certimate/pkg/core/deployer/providers/baotapanelgo-console"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	Registries.MustRegister(domain.DeploymentProviderTypeBaotaPanelGoConsole, func(options *ProviderFactoryOptions) (deployer.Provider, error) {
		credentials := domain.AccessConfigForBaotaPanelGo{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := baotapanelgoconsole.NewDeployer(&baotapanelgoconsole.DeployerConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiKey:                   credentials.ApiKey,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
		})
		return provider, err
	})
}
