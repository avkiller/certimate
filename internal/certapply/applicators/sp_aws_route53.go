package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	awsroute53 "github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/aws-route53"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeAWSRoute53, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForAWS{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := awsroute53.NewChallengeProvider(&awsroute53.ChallengeProviderConfig{
			AccessKeyId:           credentials.AccessKeyId,
			SecretAccessKey:       credentials.SecretAccessKey,
			Region:                xmaps.GetString(options.ProviderExtendedConfig, "region"),
			HostedZoneId:          xmaps.GetString(options.ProviderExtendedConfig, "hostedZoneId"),
			DnsPropagationTimeout: options.DnsPropagationTimeout,
			DnsTTL:                options.DnsTTL,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}

	if err := ACMEDns01Registries.RegisterAlias(domain.ACMEDns01ProviderTypeAWS, domain.ACMEDns01ProviderTypeAWSRoute53); err != nil {
		panic(err)
	}
}
