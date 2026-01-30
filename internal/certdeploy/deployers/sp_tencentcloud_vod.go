package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudvod "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-vod"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudVOD, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudvod.NewSSLDeployerProvider(&tencentcloudvod.SSLDeployerProviderConfig{
			SecretId:  credentials.SecretId,
			SecretKey: credentials.SecretKey,
			Endpoint:  xmaps.GetString(options.ProviderExtendedConfig, "endpoint"),
			SubAppId:  xmaps.GetInt64(options.ProviderExtendedConfig, "subAppId"),
			Domain:    xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
