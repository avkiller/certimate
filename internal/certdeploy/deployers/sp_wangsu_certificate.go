package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	wangsucertificate "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/wangsu-certificate"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeWangsuCertificate, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForWangsu{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := wangsucertificate.NewSSLDeployerProvider(&wangsucertificate.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			AccessKeySecret: credentials.AccessKeySecret,
			CertificateId:   xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
