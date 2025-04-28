// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package opens

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/infrastructure/client"
)

type Opens interface {
	// 新規注文
	GET(ctx context.Context) (*output.Opens, error)
}

type opens struct {
	client client.Client
}

func New(client client.Client) Opens {
	return &opens{
		client,
	}
}

func (o opens) GET(ctx context.Context) (*output.Opens, error) {
	req, err := o.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/opens",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.Opens
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
