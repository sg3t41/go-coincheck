package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/ordersrate"
)

func (rest *rest) OrdersRate(ctx context.Context, pair, orderType string, price, amount float64) (*ordersrate.GetResponse, error) {
	return rest.orders_rate.GET(ctx, pair, orderType, price, amount)
}
