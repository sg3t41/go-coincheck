package ordersrate

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type OrdersRate interface {
	GET(ctx context.Context, params input.OrdersRate) (*output.OrdersRate, error)
}

type ordersRate struct {
	client client.Client
}

func New(client client.Client) OrdersRate {
	return &ordersRate{
		client,
	}
}

func (r ordersRate) GET(ctx context.Context, params input.OrdersRate) (*output.OrdersRate, error) {
	queryParams := map[string]string{
		"order_type": params.OrderType,
		"pair":       params.Pair,
	}

	// orderTypeがsellの場合はpriceを設定
	if params.OrderType == "sell" {
		queryParams["price"] = strconv.FormatFloat(params.Price, 'f', -1, 64)
	}
	// orderTypeがbuyの場合はamountを設定
	if params.OrderType == "buy" {
		queryParams["amount"] = strconv.FormatFloat(params.Amount, 'f', -1, 64)
	}

	req, err := r.client.CreateRequest(ctx, http_client.RequestInput{
		Method:     http.MethodGet,
		Path:       "/api/exchange/orders/rate",
		QueryParam: queryParams,
	})
	if err != nil {
		return nil, err
	}

	var res output.OrdersRate
	if err := r.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
