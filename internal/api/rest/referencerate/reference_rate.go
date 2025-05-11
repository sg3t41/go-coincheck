// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package referencerate

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type ReferenceRate interface {
	GET(ctx context.Context, pair string) (*GetResponse, error)
}

type referenceRate struct {
	client client.Client
}

func New(client client.Client) ReferenceRate {
	return &referenceRate{
		client,
	}
}

type GetResponse struct {
	Rate string `json:"rate"`
}

func (t referenceRate) GET(ctx context.Context, pair string) (*GetResponse, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/rate/" + pair,
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
