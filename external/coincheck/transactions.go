package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetTransactions は、取引履歴を取得します
func (c *coincheck) Transactions(ctx context.Context) (*output.GetTransactions, error) {
	return c.transactions.GET(ctx)
}
