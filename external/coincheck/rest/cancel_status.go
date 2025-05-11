package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/cancelstatus"
)

func (rest *rest) CancelStatus(ctx context.Context, id int) (*cancelstatus.GetResponse, error) {
	return rest.cancel_status.GET(ctx, id)
}
