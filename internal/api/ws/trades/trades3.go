package trades

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const (
	wsURL = "wss://ws-api.coincheck.com/"
)

// SubscribeMessage represents the subscription request structure
type SubscribeMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

func main() {
	// 1. WebSocket接続の作成
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	// 2. Subscribeメッセージの送信
	subscribeToChannel(conn, "btc_jpy-orderbook")
	subscribeToChannel(conn, "btc_jpy-trades")

	// 3. メッセージの受信
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		fmt.Printf("Received: %s\n", message)
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
		log.Fatalf("Failed to encode subscribe message: %v", err)
	}

	// Subscribeメッセージを送信
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatalf("Failed to send subscribe message: %v", err)
	}
	fmt.Printf("Subscribed to channel: %s\n", channel)
}
