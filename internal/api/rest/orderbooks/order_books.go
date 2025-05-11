package orderbooks

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

var endpoint = "/api/order_books"

type OrderBooks interface {
	GET(ctx context.Context, pair string) (*GetResponse, error)
}

type orderBooks struct {
	client client.Client
}

func New(client client.Client) OrderBooks {
	return &orderBooks{client}
}

// GetOrderBooks はオーダーブックの情報を保持する
type GetResponse struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

func (t orderBooks) GET(ctx context.Context, pair string) (*GetResponse, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   endpoint,
		QueryParam: map[string]string{
			"pair": pair,
		},
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
