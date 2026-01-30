package applicators

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"

	"github.com/certimate-go/certimate/internal/domain"
	aliyunesa "github.com/certimate-go/certimate/pkg/core/ssl-applicator/acme-dns01/providers/aliyun-esa"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := ACMEDns01Registries.Register(domain.ACMEDns01ProviderTypeAliyunESA, func(options *ProviderFactoryOptions) (challenge.Provider, error) {
		credentials := domain.AccessConfigForAliyun{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := aliyunesa.NewChallengeProvider(&aliyunesa.ChallengeProviderConfig{
			AccessKeyId:           credentials.AccessKeyId,
			AccessKeySecret:       credentials.AccessKeySecret,
			Region:                xmaps.GetString(options.ProviderExtendedConfig, "region"),
			DnsPropagationTimeout: options.DnsPropagationTimeout,
			DnsTTL:                options.DnsTTL,
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
