package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetTicker は、指定されたペアのティッカー情報を取得します
func (c *coincheck) Ticker(ctx context.Context, in input.GetTicker) (*output.GetTicker, error) {
	return c.ticker.GET(ctx, in)
}
