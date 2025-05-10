// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package ticker

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Ticker interface {
	GET(context.Context, input.GetTicker) (*output.GetTicker, error)
}

type ticker struct {
	client client.Client
}

func New(client client.Client) Ticker {
	return &ticker{
		client,
	}
}

// GetTicker は最新のティッカー情報を取得する
// API: GET /api/ticker
// 可視性: パブリック
// Pair を指定しない場合は、デフォルトで btc_jpy の情報を取得する
func (t ticker) GET(ctx context.Context, in input.GetTicker) (*output.GetTicker, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/ticker",
		QueryParam: map[string]string{
			"pair": in.Pair,
		},
	})
	if err != nil {
		return nil, err
	}

	var res output.GetTicker
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
