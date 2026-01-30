package deployers

import (
	"fmt"
	"net/http"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	webhook "github.com/certimate-go/certimate/pkg/core/ssl-deployer/providers/webhook"
	xhttp "github.com/certimate-go/certimate/pkg/utils/http"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.DeploymentProviderTypeWebhook, func(options *ProviderFactoryOptions) (core.SSLDeployer, error) {
		credentials := domain.AccessConfigForWebhook{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		mergedHeaders := make(map[string]string)
		if defaultHeadersString := credentials.HeadersString; defaultHeadersString != "" {
			h, err := xhttp.ParseHeaders(defaultHeadersString)
			if err != nil {
				return nil, fmt.Errorf("failed to parse webhook headers: %w", err)
			}
			for key := range h {
				mergedHeaders[http.CanonicalHeaderKey(key)] = h.Get(key)
			}
		}
		if extendedHeadersString := xmaps.GetString(options.ProviderExtendedConfig, "headers"); extendedHeadersString != "" {
			h, err := xhttp.ParseHeaders(extendedHeadersString)
			if err != nil {
				return nil, fmt.Errorf("failed to parse webhook headers: %w", err)
			}
			for key := range h {
				mergedHeaders[http.CanonicalHeaderKey(key)] = h.Get(key)
			}
		}

		provider, err := webhook.NewSSLDeployerProvider(&webhook.SSLDeployerProviderConfig{
			WebhookUrl:               credentials.Url,
			WebhookData:              xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "webhookData", credentials.DataString),
			Method:                   credentials.Method,
			Headers:                  mergedHeaders,
			Timeout:                  xmaps.GetInt32(options.ProviderExtendedConfig, "timeout"),
			AllowInsecureConnections: credentials.AllowInsecureConnections,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
