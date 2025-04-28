package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// CancelOrder は、指定された注文をキャンセルします
func (c *coincheck) CancelOrder(ctx context.Context, in input.CancelOrder) (*output.CancelOrder, error) {
	return c.orders.DELETE(ctx, in)
}
