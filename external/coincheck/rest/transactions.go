package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactions"
)

func (c *rest) Transactions(ctx context.Context) (*transactions.GetReponse, error) {
	return c.transactions.GET(ctx)
}
