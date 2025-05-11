// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package cancelstatus

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type CancelStatus interface {
	// 新規注文
	GET(ctx context.Context, id int) (*GetResponse, error)
}

type cancelStatus struct {
	client client.Client
}

func New(client client.Client) CancelStatus {
	return &cancelStatus{
		client,
	}
}

// CancelOrder は注文キャンセルのレスポンス
type GetResponse struct {
	Success   bool   `json:"success"`    // 成功フラグ
	ID        int    `json:"id"`         // 注文のID
	Cancel    bool   `json:"cancel"`     // キャンセル済みか（true or false）
	CreatedAt string `json:"created_at"` // 注文の作成日時
}

func (c cancelStatus) GET(ctx context.Context, id int) (*GetResponse, error) {
	req, err := c.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/cancel_status?id=" + strconv.Itoa(id),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := c.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
