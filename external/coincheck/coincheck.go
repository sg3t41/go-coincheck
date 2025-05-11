package coincheck

import (
	"github.com/sg3t41/go-coincheck/external/coincheck/rest"
	"github.com/sg3t41/go-coincheck/external/coincheck/ws"
)

type Coincheck struct {
	REST rest.REST
	WS   ws.WS
}

func New(key, secret string) (*Coincheck, error) {
	rest, err := rest.New(key, secret)
	if err != nil {
		return nil, err
	}

	ws, err := ws.New() // TODO 認証情報optional化
	if err != nil {
		return nil, err
	}

	return &Coincheck{
		REST: rest,
		WS:   ws,
	}, nil
}
