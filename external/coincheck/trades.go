package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetTrades は、指定されたペアの取引履歴を取得します
func (c *coincheck) Trades(ctx context.Context, in input.GetTrades) (*output.GetTrades, error) {
	return c.trades.GET(ctx, in)
}
