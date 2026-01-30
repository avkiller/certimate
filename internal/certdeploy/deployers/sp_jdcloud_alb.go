package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	jdcloudalb "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/jdcloud-alb"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeJDCloudALB, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForJDCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := jdcloudalb.NewSSLDeployerProvider(&jdcloudalb.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			RegionId:        xmaps.GetString(options.ProviderExtendedConfig, "regionId"),
			ResourceType:    jdcloudalb.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			LoadbalancerId:  xmaps.GetString(options.ProviderExtendedConfig, "loadbalancerId"),
			ListenerId:      xmaps.GetString(options.ProviderExtendedConfig, "listenerId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
