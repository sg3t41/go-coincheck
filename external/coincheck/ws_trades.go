package coincheck

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	wsURL = "wss://ws-api.coincheck.com/"
)

type SubscribeMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

func (c *coincheck) WebSocketTrade(ctx context.Context, channel string, tradeChan chan<- string) error {
	for {
		log.Println("[INFO] Starting WebSocket connection to:", wsURL)

		// WebSocket接続の作成
		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			log.Printf("[ERROR] Failed to connect to WebSocket server: %v", err)
			time.Sleep(5 * time.Second) // 再接続を試みる前に待機
			continue
		}

		log.Println("[INFO] WebSocket connection established")

		// 切断時に接続を閉じる
		go func() {
			<-ctx.Done()
			log.Println("[INFO] Context cancelled, closing WebSocket connection")
			conn.Close()
		}()

		// Subscribeメッセージの送信
		subscribeToChannel(conn, channel)
		//		subscribeToChannel(conn, "btc_jpy-orderbook")

		// メッセージの受信ループ
		err = readMessages(ctx, conn, tradeChan)
		if err != nil {
			log.Printf("[ERROR] Error reading message: %v", err)
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

func readMessages(ctx context.Context, conn *websocket.Conn, tradeChan chan<- string) error {
	log.Println("[INFO] Starting to read messages")
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[ERROR] Failed to read message: %v", err)
			return err
		}

		//		log.Printf("[DEBUG] Received message: %s", string(message))

		// メッセージをチャネルに送信
		select {
		case tradeChan <- string(message):
			//			log.Println("[INFO] Message sent to tradeChan")
		case <-ctx.Done():
			log.Println("[INFO] Context cancelled in readMessages, exiting")
			return nil
		default:
			log.Println("[WARN] tradeChan is full, skipping message")
		}
	}
}

func subscribeToChannel(conn *websocket.Conn, channel string) {
	// Subscribeメッセージを構築
	subscribeMessage := SubscribeMessage{
		Type:    "subscribe",
		Channel: channel,
	}

	message, err := json.Marshal(subscribeMessage)
	if err != nil {
		log.Fatalf("[FATAL] Failed to encode subscribe message: %v", err)
	}

	// Subscribeメッセージを送信
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatalf("[FATAL] Failed to send subscribe message: %v", err)
	}
	log.Printf("[INFO] Subscribed to channel: %s", channel)
}
