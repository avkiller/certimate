package namesilo

import (
	"errors"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/namesilo"
)

type NameSiloApplicantConfig struct {
	ApiKey                string `json:"apiKey"`
	DnsPropagationTimeout int32  `json:"dnsPropagationTimeout,omitempty"`
	DnsTTL                int32  `json:"dnsTTL,omitempty"`
}

func NewChallengeProvider(config *NameSiloApplicantConfig) (challenge.Provider, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	providerConfig := namesilo.NewDefaultConfig()
	providerConfig.APIKey = config.ApiKey
	if config.DnsPropagationTimeout != 0 {
		providerConfig.PropagationTimeout = time.Duration(config.DnsPropagationTimeout) * time.Second
	}
	if config.DnsTTL != 0 {
		providerConfig.TTL = int(config.DnsTTL)
	}

	provider, err := namesilo.NewDNSProviderConfig(providerConfig)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
