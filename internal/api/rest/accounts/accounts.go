package accounts

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"

	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Accounts interface {
	Get(context.Context) (*Response, error)
}

type accounts struct {
	client client.Client
}

func New(client client.Client) Accounts {
	return &accounts{
		client: client,
	}
}

// Accounts はアカウント情報を持つ構造体
type Response struct {
	Success        bool                       `json:"success"`
	ID             int                        `json:"id"`
	Email          string                     `json:"email"`
	IdentityStatus string                     `json:"identity_status"`
	BitcoinAddress string                     `json:"bitcoin_address"`
	TakerFee       string                     `json:"taker_fee"`
	MakerFee       string                     `json:"maker_fee"`
	ExchangeFees   map[string]ExchangeFeeRate `json:"exchange_fees"`
}

// ExchangeFeeRate は板ごとの手数料情報を持つ構造体
type ExchangeFeeRate struct {
	MakerFeeRate string `json:"maker_fee_rate"`
	TakerFeeRate string `json:"taker_fee_rate"`
}

func (accounts accounts) Get(context context.Context) (*Response, error) {
	req, err := accounts.client.CreateRequest(context, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res Response
	if err := accounts.client.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
