package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// CreateOrder は、新しい注文を作成します
func (c *coincheck) CreateOrder(ctx context.Context, i input.CreateOrder) (*output.CreateOrder, error) {
	return c.orders.POST(ctx, i)
}
