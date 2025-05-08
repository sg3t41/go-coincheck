package balance

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
)

type Balance interface {
	GET(context.Context) (*output.Balance, error)
}

type balance struct {
	client client.Client
}

func New(client client.Client) Balance {
	return &balance{client}
}

func (t balance) GET(ctx context.Context) (*output.Balance, error) {
	req, err := t.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts/balance",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.Balance
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
