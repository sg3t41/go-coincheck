package ws

import (
	"context"
)

func (w *ws) Trades(ctx context.Context, channel string, tradeChan chan<- string) error {
	return w.trades.Subscribe(ctx, channel, tradeChan)
}
