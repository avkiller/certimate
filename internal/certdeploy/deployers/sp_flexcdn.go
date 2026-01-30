package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/flexcdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeFlexCDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForFlexCDN{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := flexcdn.NewSSLDeployerProvider(&flexcdn.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiRole:                  credentials.ApiRole,
			AccessKeyId:              credentials.AccessKeyId,
			AccessKey:                credentials.AccessKey,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             flexcdn.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			CertificateId:            xmaps.GetInt64(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
