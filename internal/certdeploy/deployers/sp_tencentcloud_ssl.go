package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudssl "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-ssl"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudSSL, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudssl.NewSSLDeployerProvider(&tencentcloudssl.SSLDeployerProviderConfig{
			SecretId:  credentials.SecretId,
			SecretKey: credentials.SecretKey,
			Endpoint:  xmaps.GetString(options.ProviderExtendedConfig, "endpoint"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
