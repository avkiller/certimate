package deployers

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	tencentcloudsslupdate "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/tencentcloud-ssl-update"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeTencentCloudSSLUpdate, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForTencentCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := tencentcloudsslupdate.NewSSLDeployerProvider(&tencentcloudsslupdate.SSLDeployerProviderConfig{
			SecretId:        credentials.SecretId,
			SecretKey:       credentials.SecretKey,
			Endpoint:        xmaps.GetString(options.ProviderExtendedConfig, "endpoint"),
			CertificateId:   xmaps.GetString(options.ProviderExtendedConfig, "certificateId"),
			IsReplaced:      xmaps.GetBool(options.ProviderExtendedConfig, "isReplaced"),
			ResourceTypes:   lo.Filter(strings.Split(xmaps.GetString(options.ProviderExtendedConfig, "resourceTypes"), ";"), func(s string, _ int) bool { return s != "" }),
			ResourceRegions: lo.Filter(strings.Split(xmaps.GetString(options.ProviderExtendedConfig, "resourceRegions"), ";"), func(s string, _ int) bool { return s != "" }),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
