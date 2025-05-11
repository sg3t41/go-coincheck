package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/orderbooks"
)

func (c *rest) OrderBooks(ctx context.Context, pair string) (*orderbooks.GetResponse, error) {
	return c.order_books.GET(ctx, pair)
}
