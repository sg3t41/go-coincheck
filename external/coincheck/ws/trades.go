package ws

import (
	"context"
)

func (ws *ws) Trades(ctx context.Context, pair string) (<-chan string, error) {
	return ws.trades.Subscribe(ctx, pair)
}
