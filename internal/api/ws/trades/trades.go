package trades

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/client"
)

type Trades interface {
	SubscribeTrades(ctx context.Context) error
}

type trades struct {
	client client.Client
}

func New(client client.Client) Trades {
	return &trades{
		client: client,
	}
}

func (t *trades) SubscribeTrades(ctx context.Context) error {
	// サブスクライブするチャンネル
	channel := "btc_jpy-trades"

	t.client.Connect(ctx)
	t.client.Subscribe(channel)

	return nil
}
