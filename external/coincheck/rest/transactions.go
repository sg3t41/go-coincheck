package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactions"
)

func (rest *rest) Transactions(ctx context.Context) (*transactions.GetReponse, error) {
	return rest.transactions.GET(ctx)
}
