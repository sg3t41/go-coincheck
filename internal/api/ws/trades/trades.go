package trades

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sg3t41/go-coincheck/internal/client"
)

type WSTrades interface {
	SubscribeTrades(ctx context.Context) error
}

type trades struct {
	client client.Client
}

func New(client client.Client) WSTrades {
	return &trades{
		client: client,
	}
}

func (t *trades) SubscribeTrades(ctx context.Context) error {
	// サブスクライブするチャンネル
	channel := "btc_jpy-trades"

	t.client.Connect(ctx)
	t.client.Subscribe(channel)

	// メッセージの読み取り開始
	go func() {
		err := t.client.ReadMessages(ctx, func(message []byte) {
			fmt.Println("Received:", string(message))
		})
		if err != nil {
			log.Println("read error:", err)
		}
	}()

	// 適当に5秒待って終了（本番では select{} で待ち続けるなど）
	select {
	case <-ctx.Done():
		fmt.Println("interrupted, shutting down...")
	case <-time.After(10 * time.Second):
		fmt.Println("timeout, shutting down...")
	}

	// 終了処理
	if err := t.client.Close(); err != nil {
		log.Println("close error:", err)
	}

	return nil
}
