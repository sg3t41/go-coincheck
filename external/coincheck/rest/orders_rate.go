package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) OrdersRate(ctx context.Context, input input.OrdersRate) (*output.OrdersRate, error) {
	return c.orders_rate.GET(ctx, input)
}
