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
	Subscribe(channel string) error
	ReadMessages(context.Context, func([]byte)) error
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

func (c *webSocketClient) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Create a WebSocket connection
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, c.url, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	log.Println("WebSocket connected to", c.url)
	return nil
}

func (c *webSocketClient) Subscribe(channel string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subscribed[channel] {
		return nil // Already subscribed
	}

	subscribeMessage := map[string]string{
		"type":    "subscribe",
		"channel": channel,
	}
	message, err := json.Marshal(subscribeMessage)
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}
	c.subscribed[channel] = true
	log.Println("Subscribed to channel:", channel)
	return nil
}

func (c *webSocketClient) ReadMessages(ctx context.Context, handler func([]byte)) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(c.readTimeout))
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}
		go handler(message)
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
