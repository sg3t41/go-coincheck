package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/trades"
)

func (c *rest) Trades(ctx context.Context, pair string) (*trades.GetResponse, error) {
	return c.trades.GET(ctx, pair)
}
