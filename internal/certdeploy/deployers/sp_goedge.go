package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/goedge"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeGoEdge, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForGoEdge{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := goedge.NewSSLDeployerProvider(&goedge.SSLDeployerProviderConfig{
			ServerUrl:                credentials.ServerUrl,
			ApiRole:                  credentials.ApiRole,
			AccessKeyId:              credentials.AccessKeyId,
			AccessKey:                credentials.AccessKey,
			AllowInsecureConnections: credentials.AllowInsecureConnections,
			ResourceType:             goedge.ResourceType(xmaps.GetString(options.ProviderExtendedConfig, "resourceType")),
			CertificateId:            xmaps.GetInt64(options.ProviderExtendedConfig, "certificateId"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
