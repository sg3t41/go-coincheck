package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts/balance"
)

func (rest *rest) Balance(ctx context.Context) (*balance.GetReponse, error) {
	return rest.balance.GET(ctx)
}
