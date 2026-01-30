package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/gname"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeGname, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForGname{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := gname.NewChallengeProvider(&gname.ChallengeProviderConfig{
			AppId:                 credentials.AppId,
			AppKey:                credentials.AppKey,
			DnsPropagationTimeout: options.DnsPropagationTimeout,
			DnsTTL:                options.DnsTTL,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
