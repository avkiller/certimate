package notify

import (
	"context"

	"github.com/certimate-go/certimate/internal/domain"
)

type accessRepository interface {
	GetById(ctx context.Context, id string) (*domain.Access, error)
}
