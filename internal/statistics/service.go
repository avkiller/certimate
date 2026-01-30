package statistics

import (
	"context"

	"github.com/certimate-go/certimate/internal/domain"
)

type StatisticsService struct {
	statRepo statisticsRepository
}

func NewStatisticsService(statRepo statisticsRepository) *StatisticsService {
	return &StatisticsService{
		statRepo: statRepo,
	}
}

func (s *StatisticsService) Get(ctx context.Context) (*domain.Statistics, error) {
	return s.statRepo.Get(ctx)
}
