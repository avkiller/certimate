package deployers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	azurekeyvault "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/azure-keyvault"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeAzureKeyVault, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForAzure{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := azurekeyvault.NewSSLDeployerProvider(&azurekeyvault.SSLDeployerProviderConfig{
			TenantId:        credentials.TenantId,
			ClientId:        credentials.ClientId,
			ClientSecret:    credentials.ClientSecret,
			CloudName:       credentials.CloudName,
			KeyVaultName:    xmaps.GetString(options.ProviderExtendedConfig, "keyvaultName"),
			CertificateName: xmaps.GetString(options.ProviderExtendedConfig, "certificateName"),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
