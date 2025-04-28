package trades

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/infrastructure/client"
)

type Trades interface {
	GET(context.Context, input.GetTrades) (*output.GetTrades, error)
}

type trades struct {
	client client.Client
}

func New(client client.Client) Trades {
	return &trades{
		client,
	}
}

func (t trades) GET(ctx context.Context, in input.GetTrades) (*output.GetTrades, error) {
	req, err := t.client.CreateRequest(ctx, client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/trades",
		QueryParam: map[string]string{
			"pair": in.Pair,
		},
	})
	if err != nil {
		return nil, err
	}

	var res output.GetTrades
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
