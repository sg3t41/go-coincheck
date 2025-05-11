package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *coincheck) Transactions(ctx context.Context) (*output.GetTransactions, error) {
	return c.transactions.GET(ctx)
}
