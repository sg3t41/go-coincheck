package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchangestatus"
)

func (rest *rest) ExchangeStatus(ctx context.Context, pair string) (*exchangestatus.GetReponse, error) {
	return rest.exchange_status.GET(ctx, pair)
}
