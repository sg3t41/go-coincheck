package orderbooks

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
)

var endpoint = "/api/order_books"

type OrderBooks interface {
	GET(context.Context, input.GetOrderBooks) (*output.GetOrderBooks, error)
}

type orderBooks struct {
	client client.Client
}

func New(client client.Client) OrderBooks {
	return &orderBooks{client}
}

func (t orderBooks) GET(ctx context.Context, in input.GetOrderBooks) (*output.GetOrderBooks, error) {
	req, err := t.client.CreateRequest(ctx, client.RequestInput{
		Method: http.MethodGet,
		Path:   endpoint,
		QueryParam: map[string]string{
			"pair": in.Pair,
		},
	})
	if err != nil {
		return nil, err
	}

	var res output.GetOrderBooks
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
