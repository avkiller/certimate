package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	aliyunga "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/aliyun-ga"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAliyunGA, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAliyun{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := aliyunga.NewSSLDeployerProvider(&aliyunga.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			ResourceGroupId: credentials.ResourceGroupId,
			ResourceType:    aliyunga.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			AcceleratorId:   xmaps.GetString(options.ProviderExtendedConfig, "acceleratorId"),
			ListenerId:      xmaps.GetString(options.ProviderExtendedConfig, "listenerId"),
			Domain:          xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
