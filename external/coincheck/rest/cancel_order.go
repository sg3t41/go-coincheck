package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) CancelOrder(ctx context.Context, in input.CancelOrder) (*output.CancelOrder, error) {
	return c.orders.DELETE(ctx, in)
}
