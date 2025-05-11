package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) Ticker(ctx context.Context, in input.GetTicker) (*output.GetTicker, error) {
	return c.ticker.GET(ctx, in)
}
