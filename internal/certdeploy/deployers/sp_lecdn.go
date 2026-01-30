package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/lecdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeLeCDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForLeCDN{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := lecdn.NewSSLDeployerProvider(&lecdn.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiVersion:               credentials.ApiVersion,
			ApiRole:                  credentials.ApiRole,
			Username:                 credentials.Username,
			Password:                 credentials.Password,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             lecdn.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			CertificateId:            xmaps.GetInt64(options.ProviderExtendedConfig, "certificateId"),
			ClientId:                 xmaps.GetInt64(options.ProviderExtendedConfig, "clientId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
