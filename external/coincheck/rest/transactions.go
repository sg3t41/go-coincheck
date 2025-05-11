package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) Transactions(ctx context.Context) (*output.GetTransactions, error) {
	return c.transactions.GET(ctx)
}
