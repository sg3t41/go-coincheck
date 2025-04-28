package application

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/domain"
)

type Accounts interface {
	Get(context.Context) (*output.Accounts, error)
}

type accounts struct {
	service domain.AccountsService
}

func NewAccounts(service domain.AccountsService) Accounts {
	return &accounts{
		service: service,
	}
}

func (a accounts) Get(ctx context.Context) (*output.Accounts, error) {
	return a.service.Get(ctx)
}
