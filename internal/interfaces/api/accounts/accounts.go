package accounts

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/application"
)

type Accounts interface {
	Get(context.Context) (*output.Accounts, error)
}

type accounts struct {
	application application.Accounts
}

func New(application application.Accounts) Accounts {
	return &accounts{
		application: application,
	}
}

func (accounts accounts) Get(context context.Context) (*output.Accounts, error) {
	return accounts.application.Get(context)
}
