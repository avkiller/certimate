package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	aliyunalb "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/aliyun-alb"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAliyunALB, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAliyun{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := aliyunalb.NewSSLDeployerProvider(&aliyunalb.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			ResourceGroupId: credentials.ResourceGroupId,
			Region:          xmaps.GetString(options.ProviderExtendedConfig, "region"),
			ResourceType:    aliyunalb.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			LoadbalancerId:  xmaps.GetString(options.ProviderExtendedConfig, "loadbalancerId"),
			ListenerId:      xmaps.GetString(options.ProviderExtendedConfig, "listenerId"),
			Domain:          xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
