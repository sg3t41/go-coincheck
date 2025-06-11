package balance

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Balance interface {
	GET(context.Context) (*GetResponse, error)
}

type balance struct {
	client client.Client
}

func New(client client.Client) Balance {
	return &balance{client}
}

// Balance はアカウントの残高情報を持つ構造体
type GetResponse struct {
	Success      bool   `json:"success"`
	JPY          string `json:"jpy"`
	BTC          string `json:"btc"`
	JPYReserved  string `json:"jpy_reserved"`
	BTCReserved  string `json:"btc_reserved"`
	JPYLending   string `json:"jpy_lending"`
	BTCLending   string `json:"btc_lending"`
	JPYLendInUse string `json:"jpy_lend_in_use"`
	BTCLendInUse string `json:"btc_lend_in_use"`
	JPYLent      string `json:"jpy_lent"`
	BTCLent      string `json:"btc_lent"`
	JPYDebt      string `json:"jpy_debt"`
	BTCDebt      string `json:"btc_debt"`
	JPYTsumitate string `json:"jpy_tsumitate"`
	BTCTsumitate string `json:"btc_tsumitate"`
}

func (t balance) GET(ctx context.Context) (*GetResponse, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts/balance",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
