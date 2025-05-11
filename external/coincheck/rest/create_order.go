package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders"
)

func (rest *rest) CreateOrder(ctx context.Context, pair, orderType string, rate, amount float64) (*orders.PostResponse, error) {
	return rest.orders.POST(ctx, pair, orderType, rate, amount)
}
