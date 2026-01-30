package certifiers

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/certifier/challengers/dns01/qingcloud"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	ACMEDns01Registries.MustRegister(domain.ACMEDns01ProviderTypeQingCloudDNS, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForQingCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := qingcloud.NewChallenger(&qingcloud.ChallengerConfig{
			AccessKeyId:           credentials.AccessKeyId,
			SecretAccessKey:       credentials.SecretAccessKey,
			DnsPropagationTimeout: options.DnsPropagationTimeout,
			DnsTTL:                options.DnsTTL,
		})
		return provider, err
	})

	ACMEDns01Registries.MustRegisterAlias(domain.ACMEDns01ProviderTypeQingCloud, domain.ACMEDns01ProviderTypeQingCloudDNS)
}
