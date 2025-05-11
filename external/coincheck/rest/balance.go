package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts/balance"
)

func (c *rest) Balance(ctx context.Context) (*balance.GetReponse, error) {
	return c.balance.GET(ctx)
}
