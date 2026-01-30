package local

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/http/webroot"

	"github.com/certimate-go/certimate/pkg/core"
)

type ChallengeProviderConfig struct {
	// 网站根目录路径。
	WebRootPath string `json:"webRootPath"`
}

func NewChallengeProvider(config *ChallengeProviderConfig) (core.ACMEChallenger, error) {
	if config == nil {
		return nil, errors.New("the configuration of the acme challenge provider is nil")
	}

	provider, err := webroot.NewHTTPProvider(config.WebRootPath)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
