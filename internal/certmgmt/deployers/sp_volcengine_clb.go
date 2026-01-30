package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/deployer"
	volcengineclb "github.com/certimate-go/certimate/pkg/core/deployer/providers/volcengine-clb"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	Registries.MustRegister(domain.DeploymentProviderTypeVolcEngineCLB, func(options *ProviderFactoryOptions) (deployer.Provider, error) {
		credentials := domain.AccessConfigForVolcEngine{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := volcengineclb.NewDeployer(&volcengineclb.DeployerConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.SecretAccessKey,
			Region:          xmaps.GetString(options.ProviderExtendedConfig, "region"),
			ResourceType:    xmaps.GetString(options.ProviderExtendedConfig, "resourceType"),
			LoadbalancerId:  xmaps.GetString(options.ProviderExtendedConfig, "loadbalancerId"),
			ListenerId:      xmaps.GetString(options.ProviderExtendedConfig, "listenerId"),
		})
		return provider, err
	})
}
