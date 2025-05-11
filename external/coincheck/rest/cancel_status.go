package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/cancelstatus"
)

func (c *rest) CancelStatus(ctx context.Context, id int) (*cancelstatus.GetResponse, error) {
	return c.cancel_status.GET(ctx, id)
}
