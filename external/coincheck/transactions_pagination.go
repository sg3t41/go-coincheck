package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetTransactionsPagination は、ページネーション付きの取引履歴を取得します
func (c *coincheck) TransactionsPagination(
	ctx context.Context,
	in input.TransactionsPagination,
) (*output.TransactionsPagination, error) {

	return c.transactions_pagination.GET(ctx, in)
}
