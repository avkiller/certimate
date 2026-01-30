package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	ctcccloudlvdn "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/ctcccloud-lvdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeCTCCCloudLVDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForCTCCCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := ctcccloudlvdn.NewSSLDeployerProvider(&ctcccloudlvdn.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			SecretAccessKey: credentials.SecretAccessKey,
			Domain:          xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
