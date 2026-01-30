package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/baiducloud"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeBaiduCloudDNS, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForBaiduCloud{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := baiducloud.NewChallengeProvider(&baiducloud.ChallengeProviderConfig{
			AccessKeyId:           credentials.AccessKeyId,
			SecretAccessKey:       credentials.SecretAccessKey,
			DnsPropagationTimeout: options.DnsPropagationTimeout,
			DnsTTL:                options.DnsTTL,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}

	if err := ACMEDns01Registries.RegisterAlias(domain.ACMEDns01ProviderTypeBaiduCloud, domain.ACMEDns01ProviderTypeBaiduCloudDNS); err != nil {
		panic(err)
	}
}
