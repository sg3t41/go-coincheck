package domain

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

type BalanceService interface {
	Get(context.Context) (*output.Balance, error)
}

type balanceService struct {
	repository BalanceRepository
}

type BalanceRepository interface {
	Get(context.Context) (*output.Balance, error)
}

func NewBalanceService(repository BalanceRepository) BalanceService {
	return &balanceService{
		repository: repository,
	}
}

func (service balanceService) Get(ctx context.Context) (*output.Balance, error) {
	return service.repository.Get(ctx)
}
