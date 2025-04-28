package domain

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

type AccountsService interface {
	Get(context.Context) (*output.Accounts, error)
}

type accountsService struct {
	repository AccountsRepository
}

type AccountsRepository interface {
	Get(context.Context) (*output.Accounts, error)
}

func NewAccountsService(repository AccountsRepository) AccountsService {
	return &accountsService{
		repository: repository,
	}
}

func (service accountsService) Get(ctx context.Context) (*output.Accounts, error) {
	return service.repository.Get(ctx)
}
