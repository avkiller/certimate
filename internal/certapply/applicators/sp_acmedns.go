package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/acmedns"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeACMEDNS, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForACMEDNS{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := acmedns.NewChallengeProvider(&acmedns.ChallengeProviderConfig{
			ServerUrl:   credentials.ServerUrl,
			Credentials: credentials.Credentials,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
