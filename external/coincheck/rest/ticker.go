package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/ticker"
)

func (c *rest) Ticker(ctx context.Context, pair string) (*ticker.GetResponse, error) {
	return c.ticker.GET(ctx, pair)
}
