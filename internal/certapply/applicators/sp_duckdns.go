package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	duckdns "github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/duckdns"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeDuckDNS, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForDuckDNS{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := duckdns.NewChallengeProvider(&duckdns.ChallengeProviderConfig{
			Token:                 credentials.Token,
			DnsPropagationTimeout: options.DnsPropagationTimeout,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
