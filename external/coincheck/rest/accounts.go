package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts"
)

func (coincheck *rest) Accounts(ctx context.Context) (*accounts.Response, error) {
	return coincheck.accounts.Get(ctx)
}
