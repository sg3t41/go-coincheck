package orderbook

import (
	"context"
	"fmt"

	"github.com/sg3t41/go-coincheck/internal/client"
)

type Orderbook interface {
	Subscribe(ctx context.Context, channel string) (<-chan string, error)
}

type orderbook struct {
	client client.Client
}

func New(client client.Client) Orderbook {
	return &orderbook{client}
}

type SubscribeMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

func (o *orderbook) Subscribe(ctx context.Context, pair string) (<-chan string, error) {
	if err := o.client.Connect(ctx); err != nil {
		return nil, err
	}

	ch := make(chan string, 100)

	channel := fmt.Sprintf("%s%s", pair, "-orderbook")
	if err := o.client.Subscribe(ctx, channel, ch); err != nil {
		return nil, err
	}

	return ch, nil
}
