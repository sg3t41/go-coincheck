package ws

import (
	"context"
)

func (w *ws) Trades(ctx context.Context, pair string) (<-chan string, error) {
	return w.trades.Subscribe(ctx, pair)
}
