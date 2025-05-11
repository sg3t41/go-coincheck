package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) OrderBooks(ctx context.Context, i input.GetOrderBooks) (*output.GetOrderBooks, error) {
	return c.order_books.GET(ctx, i)
}
