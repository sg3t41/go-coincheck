package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetAccountBalance は、アカウントのバランスを取得します
func (c *coincheck) Balance(ctx context.Context) (*output.Balance, error) {
	return c.balance.GET(ctx)
}
