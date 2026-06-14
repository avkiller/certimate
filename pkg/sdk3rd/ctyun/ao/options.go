package ao

import (
	"github.com/certimate-go/certimate/pkg/sdk3rd/ctyun/openapi"
)

func WithAkSk(ak, sk string) openapi.OptionsFunc {
	return openapi.WithAkSk(ak, sk)
}
