package coincheck

import (
	"context"
)

// GetTrades は、指定されたペアの取引履歴を取得します
func (c *coincheck) WebSocketTrade(ctx context.Context) (any, error) {
	return nil, c.ws_trades.SubscribeTrades(ctx)
}
