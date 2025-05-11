package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactionspagination"
)

func (c *rest) TransactionsPagination(
	ctx context.Context, limit int, order string,
	startingAfter, endingBefore *int,
) (*transactionspagination.GetResponse, error) {

	return c.transactions_pagination.GET(ctx, limit, order, startingAfter, endingBefore)
}
