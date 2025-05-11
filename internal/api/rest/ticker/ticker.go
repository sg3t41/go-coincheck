// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package ticker

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Ticker interface {
	GET(ctx context.Context, pair string) (*GetResponse, error)
}

type ticker struct {
	client client.Client
}

func New(client client.Client) Ticker {
	return &ticker{
		client,
	}
}

// GetTicker は GetTicker のレスポンス
type GetResponse struct {
	// Last は最新の約定価格
	Last float64 `json:"last"`
	// Bid は現在の最高買い注文（最も高い買い希望価格）
	Bid float64 `json:"bid"`
	// Ask は現在の最安売り注文（最も安い売り希望価格）
	Ask float64 `json:"ask"`
	// High は過去24時間の最高価格
	High float64 `json:"high"`
	// Low は過去24時間の最安価格
	Low float64 `json:"low"`
	// Volume は過去24時間の取引量
	Volume float64 `json:"volume"`
	// Timestamp はデータ取得時のUNIXタイムスタンプ
	Timestamp float64 `json:"timestamp"`
}

// GetTicker は最新のティッカー情報を取得する
// API: GET /api/ticker
// 可視性: パブリック
// Pair を指定しない場合は、デフォルトで btc_jpy の情報を取得する
func (t ticker) GET(ctx context.Context, pair string) (*GetResponse, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/ticker",
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
