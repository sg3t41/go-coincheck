// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package orders

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Orders interface {
	// 新規注文
	POST(ctx context.Context, pair, orderType string, rate, amount float64) (*PostResponse, error)
	// 新規注文（オプションパラメータ付き）
	POSTWithOptions(ctx context.Context, params CreateOrderParams) (*PostResponse, error)
	DELETE(ctx context.Context, id int) (*DeleteResponse, error)
	GET(ctx context.Context, id int) (*GetResponse, error)
}

type orders struct {
	client client.Client
}

func New(client client.Client) Orders {
	return &orders{
		client,
	}
}

// CreateOrderParams は注文作成時のパラメータ
type CreateOrderParams struct {
	Pair            string   `json:"pair"`                       // 取引ペア (必須)
	OrderType       string   `json:"order_type"`                 // 注文タイプ (必須)
	Rate            *float64 `json:"rate,omitempty"`             // 注文レート (指値注文時は必須)
	Amount          *float64 `json:"amount,omitempty"`           // 注文量
	MarketBuyAmount *float64 `json:"market_buy_amount,omitempty"` // 成行買い注文時の日本円額
	StopLossRate    *float64 `json:"stop_loss_rate,omitempty"`   // 逆指値レート
	TimeInForce     *string  `json:"time_in_force,omitempty"`    // 注文有効期限 ("good_til_cancelled", "immediate_or_cancel", "fill_or_kill")
}

// GetOrder は注文のステータスを持つ構造体
type GetResponse struct {
	Success                 bool    `json:"success"`                    // 成功フラグ
	ID                      int     `json:"id"`                         // 注文のID
	Pair                    string  `json:"pair"`                       // 取引ペア
	Status                  string  `json:"status"`                     // 注文のステータス
	OrderType               string  `json:"order_type"`                 // 注文のタイプ
	Rate                    *string `json:"rate"`                       // 注文のレート (null の場合は成り行き注文)
	StopLossRate            *string `json:"stop_loss_rate"`             // 逆指値レート (null の場合もあり)
	MakerFeeRate            string  `json:"maker_fee_rate"`             // Makerとして注文を行った場合の手数料
	TakerFeeRate            string  `json:"taker_fee_rate"`             // Takerとして注文を行った場合の手数料
	Amount                  string  `json:"amount"`                     // 注文の量
	MarketBuyAmount         *string `json:"market_buy_amount"`          // 成行買で注文した日本円の金額 (null の場合もあり)
	ExecutedAmount          float64 `json:"executed_amount"`            // 約定した量
	ExecutedMarketBuyAmount *string `json:"executed_market_buy_amount"` // 成行買で約定した日本円の金額 (null の場合もあり)
	ExpiredType             string  `json:"expired_type"`               // 失効した理由
	PreventedMatchID        int     `json:"prevented_match_id"`         // 対当した注文のID
	ExpiredAmount           string  `json:"expired_amount"`             // 失効した量
	ExpiredMarketBuyAmount  *string `json:"expired_market_buy_amount"`  // 成行買で失効した日本円の金額 (null の場合もあり)
	TimeInForce             string  `json:"time_in_force"`              // 注文有効期間
	CreatedAt               string  `json:"created_at"`                 // 注文の作成日時
}

func (o orders) GET(ctx context.Context, id int) (*GetResponse, error) {
	req, err := o.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/" + strconv.Itoa(id),
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

// CreateOrder は CreateOrder のレスポンス
type PostResponse struct {
	// Success はリクエストの成功を示す
	Success bool `json:"success"`
	// ID は注文ID
	ID int `json:"id"`
	// Rate は注文レート
	Rate string `json:"rate"`
	// Amount は注文量
	Amount string `json:"amount"`
	// OrderType は注文タイプ
	OrderType string `json:"order_type"`
	// TimeInForce は注文の有効期間
	TimeInForce string `json:"time_in_force"`
	// StopLossRate はストップロスレート
	StopLossRate *string `json:"stop_loss_rate"`
	// Pair は通貨ペア
	Pair string `json:"pair"`
	// CreatedAt は作成日時
	CreatedAt time.Time `json:"created_at"`
}

func (o orders) POST(ctx context.Context, pair, orderType string, rate, amount float64) (*PostResponse, error) {
	return o.POSTWithOptions(ctx, CreateOrderParams{
		Pair:      pair,
		OrderType: orderType,
		Rate:      &rate,
		Amount:    &amount,
	})
}

func (o orders) POSTWithOptions(ctx context.Context, params CreateOrderParams) (*PostResponse, error) {
	body := make(map[string]string)
	body["pair"] = params.Pair
	body["order_type"] = params.OrderType
	
	if params.Rate != nil {
		body["rate"] = strconv.FormatFloat(*params.Rate, 'f', -1, 64)
	}
	if params.Amount != nil {
		body["amount"] = strconv.FormatFloat(*params.Amount, 'f', -1, 64)
	}
	if params.MarketBuyAmount != nil {
		body["market_buy_amount"] = strconv.FormatFloat(*params.MarketBuyAmount, 'f', -1, 64)
	}
	if params.StopLossRate != nil {
		body["stop_loss_rate"] = strconv.FormatFloat(*params.StopLossRate, 'f', -1, 64)
	}
	if params.TimeInForce != nil {
		body["time_in_force"] = *params.TimeInForce
	}
	
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := o.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodPost,
		Path:    "/api/exchange/orders",
		Body:    bytes.NewBuffer(bodyJSON),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res PostResponse
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteResponse は CancelOrder のレスポンス
type DeleteResponse struct {
	Success bool `json:"success"`
	ID      int  `json:"id"`
}

func (o orders) DELETE(ctx context.Context, id int) (*DeleteResponse, error) {
	req, err := o.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodDelete,
		Path:    "/api/exchange/orders/" + strconv.Itoa(id),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
