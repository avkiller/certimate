package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/kong"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeKong, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForKong{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := kong.NewSSLDeployerProvider(&kong.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiToken:                 credentials.ApiToken,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             kong.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			Workspace:                xmaps.GetString(options.ProviderExtendedConfig, "workspace"),
			CertificateId:            xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
