package client

import (
	"net/http"
	"net/url"

	"github.com/sg3t41/go-coincheck/internal/infrastructure/e"
)

// Optionは、Coincheckクライアントを作成する際にHTTPクライアントの詳細を設定するために指定するパラメータです。
type Option func(Client) error

// WithHTTPClientは、リクエストを行うために使用されるHTTPクライアントを設定します。
func WithHTTPClient(h *http.Client) Option {
	return func(c Client) error {
		if h == nil {
			return e.ErrNilHTTPClient
		}
		c.setHTTPClient(h)
		return nil
	}
}

// WithBaseURLは、リクエストを行うためのベースURLを設定します。
func WithBaseURL(u string) Option {
	return func(c Client) error {
		url, err := url.Parse(u)
		if err != nil {
			return e.ErrInvalidBaseURL
		}
		c.setBaseURL(url)
		return nil
	}
}

// WithCredentialsは、Coincheck APIに認証するために使用される資格情報を設定します。
func WithCredentials(key, secret string) Option {
	return func(c Client) error {
		crd := &credentials{key, secret}
		c.setCredentials(crd)

		// TODO err check
		return nil
	}
}
