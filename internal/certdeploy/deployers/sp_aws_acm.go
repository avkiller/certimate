package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	awsacm "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/aws-acm"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAWSACM, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAWS{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := awsacm.NewSSLDeployerProvider(&awsacm.SSLDeployerProviderConfig{
			AccessKeyId:     credentials.AccessKeyId,
			SecretAccessKey: credentials.SecretAccessKey,
			Region:          xmaps.GetString(options.ProviderExtendedConfig, "region"),
			CertificateArn:  xmaps.GetString(options.ProviderExtendedConfig, "certificateArn"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
