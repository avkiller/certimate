package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	ratpanelsite "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/ratpanel-site"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeRatPanelSite, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForRatPanel{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := ratpanelsite.NewSSLDeployerProvider(&ratpanelsite.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			AccessTokenId:            credentials.AccessTokenId,
			AccessToken:              credentials.AccessToken,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			SiteName:                 xmaps.GetString(options.ProviderExtendedConfig, "siteName"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
