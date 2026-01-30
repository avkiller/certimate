package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	baotapanelconsole "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/baotapanel-console"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeBaotaPanelConsole, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForBaotaPanel{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := baotapanelconsole.NewSSLDeployerProvider(&baotapanelconsole.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiKey:                   credentials.ApiKey,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			AutoRestart:              xmaps.GetBool(options.ProviderExtendedConfig, "autoRestart"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
