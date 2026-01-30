package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	huaweicloudscm "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/huaweicloud-scm"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeHuaweiCloudSCM, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForHuaweiCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := huaweicloudscm.NewSSLDeployerProvider(&huaweicloudscm.SSLDeployerProviderConfig{
			AccessKeyId:         credentials.AccessKeyId,
			SecretAccessKey:     credentials.SecretAccessKey,
			EnterpriseProjectId: credentials.EnterpriseProjectId,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
