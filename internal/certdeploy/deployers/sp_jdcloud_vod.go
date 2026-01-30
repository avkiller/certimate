package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	jdcloudvod "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/jdcloud-vod"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeJDCloudVOD, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForJDCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := jdcloudvod.NewSSLDeployerProvider(&jdcloudvod.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			Domain:          xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
