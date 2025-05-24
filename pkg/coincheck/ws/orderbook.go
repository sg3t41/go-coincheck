package ws

import (
	"context"
)

func (ws *ws) Orderbook(ctx context.Context, pair string) (<-chan string, error) {
	return ws.orderbook.Subscribe(ctx, pair)
}
