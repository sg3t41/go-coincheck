package coincheck

import (
	"github.com/sg3t41/go-coincheck/external/coincheck/rest"
	"github.com/sg3t41/go-coincheck/external/coincheck/ws"
)

type Option func(*Coincheck) error

func WithCredentials(key, secret string) Option {
	return func(c *Coincheck) error {
		cred := credentials{key, secret}
		c.credentials = cred
		return nil
	}
}

func WithREST() Option {
	return func(c *Coincheck) error {
		client, err := rest.New(c.credentials.key, c.credentials.secret)
		if err != nil {
			return err
		}
		c.REST = client
		return nil
	}
}

func WithWebSocket() Option {
	return func(c *Coincheck) error {
		client, err := ws.New()
		if err != nil {
			return err
		}
		c.WS = client
		return nil
	}
}
