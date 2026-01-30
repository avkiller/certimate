package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	baishancdn "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/baishan-cdn"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeBaishanCDN, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForBaishan{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := baishancdn.NewSSLDeployerProvider(&baishancdn.SSLDeployerProviderConfig{
			ApiToken:      credentials.ApiToken,
			Domain:        xmaps.GetString(options.ProviderExtendedConfig, "domain"),
			CertificateId: xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
