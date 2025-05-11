package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders"
)

func (c *rest) GetOrder(ctx context.Context, id int) (*orders.GetResponse, error) {
	return c.orders.GET(ctx, id)
}
