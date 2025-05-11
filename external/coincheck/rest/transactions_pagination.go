package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) TransactionsPagination(
	ctx context.Context,
	in input.TransactionsPagination,
) (*output.TransactionsPagination, error) {

	return c.transactions_pagination.GET(ctx, in)
}
