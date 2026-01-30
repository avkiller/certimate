package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/safeline"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeSafeLine, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForSafeLine{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := safeline.NewSSLDeployerProvider(&safeline.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiToken:                 credentials.ApiToken,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             safeline.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			CertificateId:            xmaps.GetInt32(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
