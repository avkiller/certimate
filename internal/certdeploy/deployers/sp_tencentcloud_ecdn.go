package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudecdn "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-ecdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudECDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudecdn.NewSSLDeployerProvider(&tencentcloudecdn.SSLDeployerProviderConfig{
			SecretId:     credentials.SecretId,
			SecretKey:    credentials.SecretKey,
			Endpoint:     xmaps.GetString(options.ProviderExtendedConfig, "endpoint"),
			MatchPattern: xmaps.GetString(options.ProviderExtendedConfig, "matchPattern"),
			Domain:       xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
