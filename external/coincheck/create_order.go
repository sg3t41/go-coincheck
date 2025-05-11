package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *coincheck) CreateOrder(ctx context.Context, i input.CreateOrder) (*output.CreateOrder, error) {
	return c.orders.POST(ctx, i)
}
