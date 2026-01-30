package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	huaweicloudelb "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/huaweicloud-elb"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeHuaweiCloudELB, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForHuaweiCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := huaweicloudelb.NewSSLDeployerProvider(&huaweicloudelb.SSLDeployerProviderConfig{
			AccessKeyId:         credentials.AccessKeyId,
			SecretAccessKey:     credentials.SecretAccessKey,
			EnterpriseProjectId: credentials.EnterpriseProjectId,
			Region:              xmaps.GetString(options.ProviderExtendedConfig, "region"),
			ResourceType:        huaweicloudelb.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			CertificateId:       xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
			LoadbalancerId:      xmaps.GetString(options.ProviderExtendedConfig, "loadbalancerId"),
			ListenerId:          xmaps.GetString(options.ProviderExtendedConfig, "listenerId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
