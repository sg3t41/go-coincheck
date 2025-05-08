// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package cancelstatus

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/client"
)

type CancelStatus interface {
	// 新規注文
	GET(context.Context, input.CancelStatus) (*output.CancelStatus, error)
}

type cancelStatus struct {
	client client.Client
}

func New(client client.Client) CancelStatus {
	return &cancelStatus{
		client,
	}
}

func (c cancelStatus) GET(ctx context.Context, in input.CancelStatus) (*output.CancelStatus, error) {
	req, err := c.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/cancel_status?id=" + strconv.Itoa(in.ID),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.CancelStatus
	if err := c.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
