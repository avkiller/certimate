package deployers

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	aliyuncasdeploy "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/aliyun-cas-deploy"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAliyunCASDeploy, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAliyun{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := aliyuncasdeploy.NewSSLDeployerProvider(&aliyuncasdeploy.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			ResourceGroupId: credentials.ResourceGroupId,
			Region:          xmaps.GetString(options.ProviderExtendedConfig, "region"),
			ResourceIds:     lo.Filter(strings.Split(xmaps.GetString(options.ProviderExtendedConfig, "resourceIds"), ";"), func(s string, _ int) bool { return s != "" }),
			ContactIds:      lo.Filter(strings.Split(xmaps.GetString(options.ProviderExtendedConfig, "contactIds"), ";"), func(s string, _ int) bool { return s != "" }),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
