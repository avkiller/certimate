package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	awscloudfront "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/aws-cloudfront"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAWSCloudFront, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAWS{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := awscloudfront.NewSSLDeployerProvider(&awscloudfront.SSLDeployerProviderConfig{
			AccessKeyId:       credentials.AccessKeyId,
			SecretAccessKey:   credentials.SecretAccessKey,
			Region:            xmaps.GetString(options.ProviderExtendedConfig, "region"),
			DistributionId:    xmaps.GetString(options.ProviderExtendedConfig, "distributionId"),
			CertificateSource: xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "certificateSource", "ACM"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
