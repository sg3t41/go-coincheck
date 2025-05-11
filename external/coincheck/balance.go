package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *coincheck) Balance(ctx context.Context) (*output.Balance, error) {
	return c.balance.GET(ctx)
}
