package coincheck

import (
	"github.com/sg3t41/go-coincheck/pkg/coincheck/rest"
	"github.com/sg3t41/go-coincheck/pkg/coincheck/ws"
)

type credentials struct{ key, secret string }

type Coincheck struct {
	REST        rest.REST
	WS          ws.WS
	credentials credentials
}

func New(opts ...Option) (*Coincheck, error) {
	c := &Coincheck{}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
