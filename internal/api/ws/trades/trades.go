package trades

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sg3t41/go-coincheck/internal/client"
)

type WSTrades interface {
	Subscribe(context.Context, string) (<-chan string, error)
}

type trades struct {
	client client.Client
}

func New(client client.Client) WSTrades {
	return &trades{
		client: client,
	}
}

func (t *trades) Subscribe(ctx context.Context, channel string) (<-chan string, error) {
	// メッセージを送信するチャネルを作成
	tradeChan := make(chan string)

	// WebSocket接続を確立
	err := t.client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	// 指定されたチャンネルを購読
	if err := t.client.Subscribe(channel); err != nil {
		return nil, fmt.Errorf("failed to subscribe to channel %s: %w", channel, err)
	}

	// メッセージの読み取りを開始
	go func() {
		defer func() {
			if err := t.client.Close(); err != nil {
				log.Println("error closing WebSocket connection:", err)
			}
			close(tradeChan) // チャネルを閉じる
		}()

		// メッセージを読み取り続ける
		err := t.client.ReadMessages(ctx, func(message []byte) {
			select {
			case tradeChan <- string(message): // メッセージをチャネルに送信
			case <-ctx.Done(): // コンテキストの終了を検知
				return
			}
		})
		if err != nil {
			log.Println("read error:", err)
		}
	}()

	// Ping/Pongメカニズムの実装
	go t.handlePingPong(ctx)

	return tradeChan, nil
}

func (t *trades) handlePingPong(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := t.client.Ping(); err != nil {
				log.Println("ping error:", err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
