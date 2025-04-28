package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *coincheck) GetOrder(ctx context.Context, in input.GetOrder) (*output.GetOrder, error) {
	return c.orders.GET(ctx, in)
}
