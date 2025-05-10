package transactions

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/output"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Transactions interface {
	GET(ctx context.Context) (*output.GetTransactions, error)
}

type transactions struct {
	client client.Client
}

func New(client client.Client) Transactions {
	return &transactions{
		client,
	}
}

func (t transactions) GET(ctx context.Context) (*output.GetTransactions, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/transactions",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.GetTransactions
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
