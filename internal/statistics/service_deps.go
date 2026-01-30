package statistics

import (
	"context"

	"github.com/certimate-go/certimate/internal/domain"
)

type statisticsRepository interface {
	Get(ctx context.Context) (*domain.Statistics, error)
}
