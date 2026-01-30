package certifiers

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	duckdns "github.com/certimate-go/certimate/pkg/core/certifier/challengers/dns01/duckdns"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	ACMEDns01Registries.MustRegister(domain.ACMEDns01ProviderTypeDuckDNS, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForDuckDNS{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := duckdns.NewChallenger(&duckdns.ChallengerConfig{
			Token:                 credentials.Token,
			DnsPropagationTimeout: options.DnsPropagationTimeout,
		})
		return provider, err
	})
}
