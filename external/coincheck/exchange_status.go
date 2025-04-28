package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetOrderBooks は、指定されたペアのオーダーブックを取得します
func (c *coincheck) ExchangeStatus(ctx context.Context, i input.ExchangeStatus) (*output.ExchangeStatus, error) {
	return c.exchange_status.GET(ctx, i)
}
