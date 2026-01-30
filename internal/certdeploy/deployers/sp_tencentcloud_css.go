package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudcss "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-css"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudCSS, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudcss.NewSSLDeployerProvider(&tencentcloudcss.SSLDeployerProviderConfig{
			SecretId:  credentials.SecretId,
			SecretKey: credentials.SecretKey,
			Endpoint:  xmaps.GetString(options.ProviderExtendedConfig, "endpoint"),
			Domain:    xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
