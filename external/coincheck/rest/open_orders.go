package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/opens"
)

func (c *rest) OpenOrders(ctx context.Context) (*opens.GetResponse, error) {
	return c.opens.GET(ctx)
}
