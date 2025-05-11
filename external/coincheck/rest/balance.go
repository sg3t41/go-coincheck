package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) Balance(ctx context.Context) (*output.Balance, error) {
	return c.balance.GET(ctx)
}
