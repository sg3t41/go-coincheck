package ws

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/ws/trades"
)

type WS interface {
	Trades(context context.Context, channel string, to chan<- string) error
}

type ws struct {
	trades trades.Trades
}

func New() (WS, error) {
	return &ws{
		trades: trades.New(),
	}, nil
}
