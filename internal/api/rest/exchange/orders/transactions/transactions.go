package transactions

import (
	"context"
	"net/http"
	"time"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Transactions interface {
	GET(ctx context.Context) (*GetReponse, error)
}

type transactions struct {
	client client.Client
}

func New(client client.Client) Transactions {
	return &transactions{
		client,
	}
}

// Transaction は取引情報を表す構造体
type Transaction struct {
	ID          int       `json:"id"`           // 取引ID
	OrderID     int       `json:"order_id"`     // 注文ID
	CreatedAt   time.Time `json:"created_at"`   // 取引日時
	Funds       Funds     `json:"funds"`        // 各残高の増減分
	Pair        string    `json:"pair"`         // 取引ペア
	Rate        string    `json:"rate"`         // 約定価格
	FeeCurrency string    `json:"fee_currency"` // 手数料の通貨
	Fee         string    `json:"fee"`          // 発生した手数料
	Liquidity   string    `json:"liquidity"`    // "T" (Taker) or "M" (Maker) or "itayose" (Itayose)
	Side        string    `json:"side"`         // "sell" or "buy"
}

// Funds は各残高の増減分を表す構造体
type Funds struct {
	BTC string `json:"btc"` // ビットコインの増減分
	JPY string `json:"jpy"` // 日本円の増減分
}

// GetTransactions は取引履歴のレスポンスを表す構造体
type GetReponse struct {
	Success      bool          `json:"success"`      // リクエストの成功を示す
	Transactions []Transaction `json:"transactions"` // 取引履歴
}

func (t transactions) GET(ctx context.Context) (*GetReponse, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/transactions",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetReponse
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
