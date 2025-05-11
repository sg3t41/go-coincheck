package transactionspagination

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactions"
	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type TransactionsPagination interface {
	GET(
		ctx context.Context, limit int, order string,
		startingAfter, endingBefore *int,
	) (*GetResponse, error)
}

type transactionsPagination struct {
	client client.Client
}

func New(client client.Client) TransactionsPagination {
	return &transactionsPagination{
		client,
	}
}

// PaginationOrder はページネーション情報を表す構造体
type PaginationOrder struct {
	Limit         int    `json:"limit"`          // 取得する取引の最大数
	Order         string `json:"order"`          // 取得する順序 ("asc" または "desc")
	StartingAfter *int   `json:"starting_after"` // 指定されたIDの取引の次の取引から取得
	EndingBefore  *int   `json:"ending_before"`  // 指定されたIDの取引の前の取引まで取得
}

// TransactionsPagination はページネーション付き取引履歴のレスポンスを表す構造体
type GetResponse struct {
	Success    bool                       `json:"success"`    // リクエストの成功を示す
	Pagination PaginationOrder            `json:"pagination"` // ページネーション情報
	Data       []transactions.Transaction `json:"data"`       // 取引履歴
}

// GetPagination は取引履歴（ページネーション）を取得する関数
func (t transactionsPagination) GET(
	ctx context.Context, limit int, order string,
	startingAfter, endingBefore *int,
) (*GetResponse, error) {

	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method: http.MethodGet,
		Path:   "/api/exchange/orders/transactions_pagination",
		QueryParam: map[string]string{
			"limit": strconv.Itoa(limit),
			"order": order,
			// "starting_after": startingAfter,
			// "ending_before":  endingBefore,
		},
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
