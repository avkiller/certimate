package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	aliyundcdn "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/aliyun-dcdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAliyunDCDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAliyun{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := aliyundcdn.NewSSLDeployerProvider(&aliyundcdn.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			ResourceGroupId: credentials.ResourceGroupId,
			Region:          xmaps.GetString(options.ProviderExtendedConfig, "region"),
			Domain:          xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
