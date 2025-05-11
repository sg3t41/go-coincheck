package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) GetOrder(ctx context.Context, in input.GetOrder) (*output.GetOrder, error) {
	return c.orders.GET(ctx, in)
}
