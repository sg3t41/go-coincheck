package repository

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/domain"
	"github.com/sg3t41/go-coincheck/internal/infrastructure/client"
)

type accounts struct {
	client client.Client
}

func NewAccounts(client client.Client) domain.AccountsRepository {
	return &accounts{
		client: client,
	}
}

func (repo accounts) Get(ctx context.Context) (*output.Accounts, error) {
	req, err := repo.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.Accounts
	if err := repo.client.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
