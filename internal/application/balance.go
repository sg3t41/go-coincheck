package application

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/domain"
)

type Balance interface {
	GET(context.Context) (*output.Balance, error)
}

type balance struct {
	service domain.BalanceService
}

func NewBalance(service domain.BalanceService) Balance {
	return &balance{
		service: service,
	}
}

func (b balance) GET(ctx context.Context) (*output.Balance, error) {
	return b.service.Get(ctx)
}
