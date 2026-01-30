package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	qiniupili "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/qiniu-pili"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeQiniuPili, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForQiniu{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := qiniupili.NewSSLDeployerProvider(&qiniupili.SSLDeployerProviderConfig{
			AccessKey: credentials.AccessKey,
			SecretKey: credentials.SecretKey,
			Hub:       xmaps.GetString(options.ProviderExtendedConfig, "hub"),
			Domain:    xmaps.GetString(options.ProviderExtendedConfig, "domain"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
