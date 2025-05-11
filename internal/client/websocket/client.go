package websocket

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketClient interface {
	Connect(ctx context.Context) error
	Subscribe(ctx context.Context, channel string, in chan<- string) error
	Close() error
}

type webSocketClient struct {
	url         string
	conn        *websocket.Conn
	mu          sync.Mutex
	subscribed  map[string]bool
	readTimeout time.Duration
}

func NewWebSocketClient() (WebSocketClient, error) {
	return &webSocketClient{
		url:         "wss://ws-api.coincheck.com/",
		subscribed:  make(map[string]bool),
		readTimeout: 10 * time.Second,
	}, nil
}

// WebSocket接続を確立する関数
func (c *webSocketClient) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	log.Println("[INFO] Starting WebSocket connection to:", c.url)
	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	c.conn = conn

	log.Println("[INFO] WebSocket connection established")

	// Contextキャンセル時に接続を閉じる
	go func() {
		<-ctx.Done()
		log.Println("[INFO] Context cancelled, closing WebSocket connection")
		c.conn.Close()
	}()

	return nil
}

func (c *webSocketClient) Subscribe(ctx context.Context, channel string, in chan<- string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	go func() {
		defer close(in)
		for {

			// Subscribeメッセージの送信
			if err := c.subscribe(ctx, channel); err != nil {
				log.Printf("[ERROR] Error subscribe message: %v\n", err)
				c.conn.Close()
				break
			}

			// メッセージの受信ループ
			if err := c.ReadMessages(ctx, in); err != nil {
				log.Printf("[ERROR] Error reading message: %v\n", err)
				c.conn.Close()
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}()

	return nil
}

type SubscribeMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

func (c *webSocketClient) subscribe(_ context.Context, channel string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

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
	err = c.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatalf("[FATAL] Failed to send subscribe message: %v", err)
	}
	log.Printf("[INFO] Subscribed to channel: %s", channel)

	return nil
}

func (c *webSocketClient) ReadMessages(ctx context.Context, in chan<- string) error {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("[ERROR] Failed to read message: %v", err)
			return err
		}

		// メッセージをチャネルに送信
		select {
		case in <- string(message):
			// メッセージを送信
		case <-ctx.Done():
			log.Println("[INFO] Context cancelled in readMessages, exiting")
			return nil
		default:
			log.Println("[WARN] tradeChan is full, skipping message")
		}
	}
}

func (c *webSocketClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return nil
	}
	err := c.conn.Close()
	log.Println("WebSocket connection closed")
	return err
}
