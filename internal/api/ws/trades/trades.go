package trades

import (
	"context"
	"fmt"

	"github.com/sg3t41/go-coincheck/internal/client"
)

type Trades interface {
	Subscribe(ctx context.Context, pair string) (<-chan string, error)
}

type trades struct {
	client client.Client
}

func New(client client.Client) Trades {
	return &trades{client}
}

type SubscribeMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

func (t *trades) Subscribe(ctx context.Context, pair string) (<-chan string, error) {
	if err := t.client.Connect(ctx); err != nil {
		return nil, err
	}

	ch := make(chan string, 100)

	channel := fmt.Sprintf("%s%s", pair, "-trades")

	if err := t.client.Subscribe(ctx, channel, ch); err != nil {
		return nil, err
	}

	return ch, nil
}
