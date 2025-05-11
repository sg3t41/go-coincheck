// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package opens

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Opens interface {
	// 新規注文
	GET(ctx context.Context) (*GetResponse, error)
}

type opens struct {
	client client.Client
}

func New(client client.Client) Opens {
	return &opens{
		client,
	}
}

// Opens は未決済注文のレスポンス
type GetResponse struct {
	Success bool    `json:"success"`
	Orders  []Order `json:"orders"`
}

// Order は未決済注文の情報
type Order struct {
	ID                     int     `json:"id"`
	OrderType              string  `json:"order_type"`
	Rate                   *string `json:"rate"` // null の場合があるためポインタ型
	Pair                   string  `json:"pair"`
	PendingAmount          *string `json:"pending_amount"`
	PendingMarketBuyAmount *string `json:"pending_market_buy_amount"`
	StopLossRate           *string `json:"stop_loss_rate"`
	CreatedAt              string  `json:"created_at"`
}

func (o opens) GET(ctx context.Context) (*GetResponse, error) {
	req, err := o.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/opens",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
