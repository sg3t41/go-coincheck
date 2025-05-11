package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders"
)

func (rest *rest) CancelOrder(ctx context.Context, id int) (*orders.DeleteResponse, error) {
	return rest.orders.DELETE(ctx, id)
}
