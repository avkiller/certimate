package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudcos "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-cos"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudCOS, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudcos.NewSSLDeployerProvider(&tencentcloudcos.SSLDeployerProviderConfig{
			SecretId:  credentials.SecretId,
			SecretKey: credentials.SecretKey,
			Region:    xmaps.GetString(options.ProviderExtendedConfig, "region"),
			Bucket:    xmaps.GetString(options.ProviderExtendedConfig, "bucket"),
			Domain:    xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
