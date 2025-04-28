package accounts

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/infrastructure/client"
)

type Accounts interface {
	Get(context.Context) (*output.Accounts, error)
}

type accounts struct {
	client client.Client
}

func New(client client.Client) Accounts {
	return &accounts{
		client: client,
	}
}

func (accounts accounts) Get(context context.Context) (*output.Accounts, error) {
	req, err := accounts.client.CreateRequest(context, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.Accounts
	if err := accounts.client.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
