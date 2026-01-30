package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	huaweicloudcdn "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/huaweicloud-cdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeHuaweiCloudCDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForHuaweiCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := huaweicloudcdn.NewSSLDeployerProvider(&huaweicloudcdn.SSLDeployerProviderConfig{
			AccessKeyId:         credentials.AccessKeyId,
			SecretAccessKey:     credentials.SecretAccessKey,
			EnterpriseProjectId: credentials.EnterpriseProjectId,
			Region:              xmaps.GetString(options.ProviderExtendedConfig, "region"),
			Domain:              xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
