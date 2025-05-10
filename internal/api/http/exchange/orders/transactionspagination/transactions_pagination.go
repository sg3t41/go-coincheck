package transactionspagination

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type TransactionsPagination interface {
	GET(ctx context.Context, i input.TransactionsPagination) (*output.TransactionsPagination, error)
}

type transactionsPagination struct {
	client client.Client
}

func New(client client.Client) TransactionsPagination {
	return &transactionsPagination{
		client,
	}
}

// GetPagination は取引履歴（ページネーション）を取得する関数
func (t transactionsPagination) GET(ctx context.Context, i input.TransactionsPagination) (*output.TransactionsPagination, error) {
	// リクエストを作成
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/exchange/orders/transactions_pagination",
		QueryParam: map[string]string{
			"limit": strconv.Itoa(i.Limit),
			"order": i.Order,
			// "starting_after": "10",
			// "ending_before":  "21",
		},
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	// レスポンスを処理
	var res output.TransactionsPagination
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
