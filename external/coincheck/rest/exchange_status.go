package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchangestatus"
)

func (c *rest) ExchangeStatus(ctx context.Context, pair string) (*exchangestatus.GetReponse, error) {
	return c.exchange_status.GET(ctx, pair)
}
