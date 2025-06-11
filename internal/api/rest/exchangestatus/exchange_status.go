package exchangestatus

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

var endpoint = "/api/exchange_status"

type ExchangeStatus interface {
	GET(ctx context.Context, pair string) (*GetResponse, error)
}

type exchangeStatus struct {
	client client.Client
}

func New(client client.Client) ExchangeStatus {
	return &exchangeStatus{client}
}

// GetExchangeStatus は取引所のステータスのレスポンス
type GetExchangeStatus struct {
	ExchangeStatus []ExchangeStatus `json:"exchange_status"`
}

// ExchangeStatus は取引所のステータス情報
type GetResponse struct {
	Pair         string       `json:"pair"`
	Status       string       `json:"status"`
	Timestamp    int64        `json:"timestamp"`
	Availability Availability `json:"availability"`
}

// Availability は注文の可用性情報
type Availability struct {
	Order       bool `json:"order"`
	MarketOrder bool `json:"market_order"`
	Cancel      bool `json:"cancel"`
}

func (es exchangeStatus) GET(ctx context.Context, pair string) (*GetResponse, error) {
	req, err := es.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   endpoint,
		QueryParam: map[string]string{
			"pair": pair,
		},
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := es.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
