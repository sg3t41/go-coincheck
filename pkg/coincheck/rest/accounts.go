package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts"
)

func (rest *rest) Accounts(ctx context.Context) (*accounts.GetResponse, error) {
	return rest.accounts.Get(ctx)
}
