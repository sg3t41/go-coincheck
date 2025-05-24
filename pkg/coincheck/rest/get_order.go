package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders"
)

func (rest *rest) GetOrder(ctx context.Context, id int) (*orders.GetResponse, error) {
	return rest.orders.GET(ctx, id)
}
