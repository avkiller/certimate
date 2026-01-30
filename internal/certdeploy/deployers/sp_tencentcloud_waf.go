package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudwaf "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-waf"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudWAF, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudwaf.NewSSLDeployerProvider(&tencentcloudwaf.SSLDeployerProviderConfig{
			SecretId:   credentials.SecretId,
			SecretKey:  credentials.SecretKey,
			Endpoint:   xmaps.GetString(options.ProviderExtendedConfig, "endpoint"),
			Domain:     xmaps.GetString(options.ProviderExtendedConfig, "domain"),
			DomainId:   xmaps.GetString(options.ProviderExtendedConfig, "domainId"),
			InstanceId: xmaps.GetString(options.ProviderExtendedConfig, "instanceId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
