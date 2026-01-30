package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/cmcccloud"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeCMCCCloudDNS, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForCMCCCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := cmcccloud.NewChallengeProvider(&cmcccloud.ChallengeProviderConfig{
			AccessKeyId:           credentials.AccessKeyId,
			AccessKeySecret:       credentials.AccessKeySecret,
			DnsPropagationTimeout: options.DnsPropagationTimeout,
			DnsTTL:                options.DnsTTL,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}

	if err := ACMEDns01Registries.RegisterAlias(domain.ACMEDns01ProviderTypeCMCCCloud, domain.ACMEDns01ProviderTypeCMCCCloudDNS); err != nil {
		panic(err)
	}
}
