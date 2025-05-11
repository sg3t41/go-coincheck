package coincheck

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// GetTrades は、指定されたペアの取引履歴を取得します
// func (c *coincheck) WebSocketTrade(ctx context.Context) (any, error) {
// 	return nil, c.ws_trades.SubscribeTrades(ctx)
// }

const (
	wsURL = "wss://ws-api.coincheck.com/"
)

type SubscribeMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

func (c *coincheck) WebSocketTrade(ctx context.Context, channel string, tradeChan chan<- string) error {
	for {
		// WebSocket接続の作成
		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			log.Printf("Failed to connect to WebSocket server: %v", err)
			time.Sleep(5 * time.Second) // 再接続を試みる前に待機
			continue
		}

		// 切断時に接続を閉じる
		go func() {
			<-ctx.Done()
			conn.Close()
		}()

		// Subscribeメッセージの送信
		subscribeToChannel(conn, channel)
		subscribeToChannel(conn, "btc_jpy-orderbook")

		// メッセージの受信ループ
		err = readMessages(ctx, conn, tradeChan)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

func readMessages(ctx context.Context, conn *websocket.Conn, tradeChan chan<- string) error {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		fmt.Printf("Received message: %s\n", string(message))

		// メッセージをチャネルに送信
		select {
		case tradeChan <- string(message):
		case <-ctx.Done():
			log.Println("Context cancelled in readMessages")
			return nil
		}
	}
	return nil
}

func subscribeToChannel(conn *websocket.Conn, channel string) {
	// Subscribeメッセージを構築
	subscribeMessage := SubscribeMessage{
		Type:    "subscribe",
		Channel: channel,
	}

	message, err := json.Marshal(subscribeMessage)
	if err != nil {
		log.Fatalf("Failed to encode subscribe message: %v", err)
	}

	// Subscribeメッセージを送信
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatalf("Failed to send subscribe message: %v", err)
	}
	fmt.Printf("Subscribed to channel: %s\n", channel)
}
