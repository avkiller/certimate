package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/deployer"
	wangsucertificate "github.com/certimate-go/certimate/pkg/core/deployer/providers/wangsu-certificate"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	Registries.MustRegister(domain.DeploymentProviderTypeWangsuCertificate, func(options *ProviderFactoryOptions) (deployer.Provider, error) {
		credentials := domain.AccessConfigForWangsu{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := wangsucertificate.NewDeployer(&wangsucertificate.DeployerConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			CertificateId:   xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	})
}
