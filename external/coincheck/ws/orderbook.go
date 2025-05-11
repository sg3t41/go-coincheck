package ws

import (
	"context"
)

func (w *ws) Orderbook(ctx context.Context, pair string) (<-chan string, error) {
	return w.orderbook.Subscribe(ctx, pair)
}
