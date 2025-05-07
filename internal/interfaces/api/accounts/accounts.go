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
		client,
	}
}

func (a accounts) Get(ctx context.Context) (*output.Accounts, error) {
	req, err := a.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.Accounts
	if err := a.client.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
