package deployers

import (
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/local"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeLocal, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		provider, err := local.NewSSLDeployerProvider(&local.SSLDeployerProviderConfig{
			ShellEnv:                 local.ShellEnvType(xmaps.GetString(options.ProviderExtendedConfig, "shellEnv")),
			PreCommand:               xmaps.GetString(options.ProviderExtendedConfig, "preCommand"),
			PostCommand:              xmaps.GetString(options.ProviderExtendedConfig, "postCommand"),
			OutputFormat:             local.OutputFormatType(xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "format", string(local.OUTPUT_FORMAT_PEM))),
			OutputCertPath:           xmaps.GetString(options.ProviderExtendedConfig, "certPath"),
			OutputServerCertPath:     xmaps.GetString(options.ProviderExtendedConfig, "certPathForServerOnly"),
			OutputIntermediaCertPath: xmaps.GetString(options.ProviderExtendedConfig, "certPathForIntermediaOnly"),
			OutputKeyPath:            xmaps.GetString(options.ProviderExtendedConfig, "keyPath"),
			PfxPassword:              xmaps.GetString(options.ProviderExtendedConfig, "pfxPassword"),
			JksAlias:                 xmaps.GetString(options.ProviderExtendedConfig, "jksAlias"),
			JksKeypass:               xmaps.GetString(options.ProviderExtendedConfig, "jksKeypass"),
			JksStorepass:             xmaps.GetString(options.ProviderExtendedConfig, "jksStorepass"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
