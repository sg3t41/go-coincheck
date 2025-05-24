package websocket

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sg3t41/go-coincheck/internal/e"
)

type WebSocketClient interface {
	Connect(ctx context.Context) error
	Subscribe(ctx context.Context, channel string, in chan<- string) error
	Close() error
}

type webSocketClient struct {
	url       *url.URL
	conn      *websocket.Conn
	mu        sync.Mutex
	connected bool
}

type Option func(*webSocketClient) error

func WithBaseURL(strURL string) Option {
	return func(wsc *webSocketClient) error {
		if strURL == "" {
			return e.WithPrefixError(errors.New("WebSocketのベースURLが空です"))
		}
		url, err := url.Parse(strURL)
		if err != nil {
			return e.WithPrefixError(err)
		}

		wsc.url = url
		return nil
	}
}

func NewClient(opts ...Option) (WebSocketClient, error) {
	defaultURL, err := url.Parse("wss://ws-api.coincheck.com/")
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	c := &webSocketClient{
		url: defaultURL,
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, e.WithPrefixError(err)
		}
	}

	return c, nil
}

// WebSocket接続を確立する関数（シングルトン的な実装）
func (c *webSocketClient) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 既に接続されている場合は何もしない
	if c.connected {
		log.Println("[INFO] WebSocket is already connected")
		return nil
	}

	log.Println("[INFO] Starting WebSocket connection to:", c.url)
	conn, _, err := websocket.DefaultDialer.Dial(c.url.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	c.conn = conn
	c.connected = true

	log.Println("[INFO] WebSocket connection established")

	// Contextキャンセル時に接続を閉じる
	go func() {
		<-ctx.Done()
		log.Println("[INFO] Context cancelled, closing WebSocket connection")
		c.Close() // 明示的にCloseを呼び出す
	}()

	return nil
}

func (c *webSocketClient) Subscribe(ctx context.Context, channel string, in chan<- string) error {
	// Connectを確認
	if err := c.Connect(ctx); err != nil {
		return err
	}

	go func() {
		defer close(in)
		for {
			// Subscribeメッセージを送信
			if err := c.subscribe(ctx, channel); err != nil {
				log.Printf("[ERROR] Error subscribing to channel: %v\n", err)
				c.Close() // エラー時に接続を閉じる
				break
			}

			// メッセージの受信ループ
			if err := c.ReadMessages(ctx, in); err != nil {
				log.Printf("[ERROR] Error reading messages: %v\n", err)
				c.Close()
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}()

	return nil
}

func (c *webSocketClient) subscribe(_ context.Context, channel string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return fmt.Errorf("WebSocket is not connected")
	}

	// Subscribeメッセージを構築
	message := fmt.Sprintf(`{"type":"subscribe","channel":"%s"}`, channel)
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send subscribe message: %w", err)
	}

	log.Printf("[INFO] Subscribed to channel: %s", channel)
	return nil
}

func (c *webSocketClient) ReadMessages(ctx context.Context, in chan<- string) error {
	for {
		if c.conn == nil {
			return fmt.Errorf("WebSocket is not connected")
		}

		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("[ERROR] Failed to read message: %v", err)
			return err
		}

		// メッセージをチャネルに送信
		select {
		case in <- string(message):
		case <-ctx.Done():
			log.Println("[INFO] Context cancelled in ReadMessages, exiting")
			return nil
		default:
			log.Println("[WARN] Message channel is full, skipping message")
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
	c.conn = nil
	c.connected = false // 接続状態をリセット
	log.Println("[INFO] WebSocket connection closed")
	return err
}
