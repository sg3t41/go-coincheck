package ordersrate

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type OrdersRate interface {
	GET(ctx context.Context, pair, orderType string, price, amount float64) (*GetResponse, error)
}

type ordersRate struct {
	client client.Client
}

func New(client client.Client) OrdersRate {
	return &ordersRate{
		client,
	}
}

type GetResponse struct {
	Success bool   `json:"success"`
	Rate    string `json:"rate"`
	Price   string `json:"price"`
	Amount  string `json:"amount"`
}

func (r ordersRate) GET(ctx context.Context, pair, orderType string, price, amount float64) (*GetResponse, error) {
	queryParams := map[string]string{
		"order_type": orderType,
		"pair":       pair,
	}

	// orderTypeがsellの場合はpriceを設定
	if orderType == "sell" {
		queryParams["price"] = strconv.FormatFloat(price, 'f', -1, 64)
	}
	// orderTypeがbuyの場合はamountを設定
	if orderType == "buy" {
		queryParams["amount"] = strconv.FormatFloat(amount, 'f', -1, 64)
	}

	req, err := r.client.CreateRequest(ctx, http_client.RequestInput{
		Method:     http.MethodGet,
		Path:       "/api/exchange/orders/rate",
		QueryParam: queryParams,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := r.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
