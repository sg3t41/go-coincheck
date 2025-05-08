// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package referencerate

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
)

type ReferenceRate interface {
	GET(context.Context, input.ReferenceRate) (*output.ReferenceRate, error)
}

type referenceRate struct {
	client client.Client
}

func New(client client.Client) ReferenceRate {
	return &referenceRate{
		client,
	}
}

func (t referenceRate) GET(ctx context.Context, i input.ReferenceRate) (*output.ReferenceRate, error) {
	req, err := t.client.CreateRequest(ctx, client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/rate/" + i.Pair,
	})
	if err != nil {
		return nil, err
	}

	var res output.ReferenceRate
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
