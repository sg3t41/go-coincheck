package ws

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/ws/orderbook"
	"github.com/sg3t41/go-coincheck/internal/api/ws/trades"
	"github.com/sg3t41/go-coincheck/internal/client"
)

type WS interface {
	Trades(context context.Context, pair string) (<-chan string, error)
	Orderbook(context context.Context, pair string) (<-chan string, error)
}

type ws struct {
	trades    trades.Trades
	orderbook orderbook.Orderbook
}

func New() (WS, error) {
	c, err := client.New(
		client.WithWebSocket("wss://ws-api.coincheck.com/"), // WebSocketクライアントのみを初期化
	)
	if err != nil {
		return nil, err
	}

	return &ws{
		trades:    trades.New(c),
		orderbook: orderbook.New(c),
	}, nil
}
