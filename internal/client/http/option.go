package http

import (
	"errors"
	"net/url"

	"github.com/sg3t41/go-coincheck/internal/e"
)

type Option func(*httpClient) error

func WithCredentials(key, secret string) Option {
	return func(hc *httpClient) error {
		// TODO fix
		switch true {
		case key == "" && secret == "":
			return e.WithPrefixError(errors.New("認証情報のkeyとsecretが空です"))
		case key == "":
			return e.WithPrefixError(errors.New("認証情報のkeyが空です"))
		case secret == "":
			return e.WithPrefixError(errors.New("認証情報のsecretが空です"))
		default:
			crd := &Credentials{key, secret}
			hc.credentials = crd
			return nil
		}
	}
}

func WithBaseURL(strURL string) Option {
	return func(hc *httpClient) error {
		if strURL == "" {
			return e.WithPrefixError(errors.New("REST APIのベースURLが空です")) // TODO fix
		}

		baseURL, err := url.Parse(strURL)
		if err != nil {
			return e.WithPrefixError(err)
		}

		hc.baseURL = baseURL
		return nil
	}
}
