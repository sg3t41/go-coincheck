package exchangestatus

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

var endpoint = "/api/exchange_status"

type ExchangeStatus interface {
	GET(context.Context, input.ExchangeStatus) (*output.ExchangeStatus, error)
}

type exchangeStatus struct {
	client client.Client
}

func New(client client.Client) ExchangeStatus {
	return &exchangeStatus{client}
}

func (es exchangeStatus) GET(ctx context.Context, i input.ExchangeStatus) (*output.ExchangeStatus, error) {
	req, err := es.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   endpoint,
		QueryParam: map[string]string{
			"pair": i.Pair,
		},
	})
	if err != nil {
		return nil, err
	}

	var res output.ExchangeStatus
	if err := es.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
