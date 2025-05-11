package trades

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Trades interface {
	GET(ctx context.Context, pair string) (*GetResponse, error)
}

type trades struct {
	client client.Client
}

func New(client client.Client) Trades {
	return &trades{
		client,
	}
}

// GetTrades は Coincheck API の取引履歴を表す
type GetResponse struct {
	Success    bool       `json:"success"`
	Pagination Pagination `json:"pagination"`
	Data       []Trade    `json:"data"`
}

// Pagination はページネーション情報
type Pagination struct {
	Limit         int    `json:"limit"`
	Order         string `json:"order"`
	StartingAfter *int   `json:"starting_after"`
	EndingBefore  *int   `json:"ending_before"`
}

// Trade は取引データ
type Trade struct {
	ID        int    `json:"id"`
	Amount    string `json:"amount"`
	Rate      string `json:"rate"`
	Pair      string `json:"pair"`
	OrderType string `json:"order_type"`
	CreatedAt string `json:"created_at"`
}

func (t trades) GET(ctx context.Context, pair string) (*GetResponse, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/trades",
		QueryParam: map[string]string{
			"pair": pair,
		},
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
