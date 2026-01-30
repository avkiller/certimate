package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	ucloudus3 "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/ucloud-us3"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeUCloudUS3, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForUCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := ucloudus3.NewSSLDeployerProvider(&ucloudus3.SSLDeployerProviderConfig{
			PrivateKey: credentials.PrivateKey,
			PublicKey:  credentials.PublicKey,
			ProjectId:  credentials.ProjectId,
			Region:     xmaps.GetString(options.ProviderExtendedConfig, "region"),
			Bucket:     xmaps.GetString(options.ProviderExtendedConfig, "bucket"),
			Domain:     xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
