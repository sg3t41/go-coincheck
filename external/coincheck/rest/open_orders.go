package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/opens"
)

func (rest *rest) OpenOrders(ctx context.Context) (*opens.GetResponse, error) {
	return rest.opens.GET(ctx)
}
