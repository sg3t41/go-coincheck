package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *coincheck) OrderBooks(ctx context.Context, i input.GetOrderBooks) (*output.GetOrderBooks, error) {
	return c.order_books.GET(ctx, i)
}
