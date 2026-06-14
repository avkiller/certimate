package cdnpro

import (
	"github.com/certimate-go/certimate/pkg/sdk3rd/wangsu/openapi"
)

func WithAkSk(ak, sk string) openapi.OptionsFunc {
	return openapi.WithAkSk(ak, sk)
}
