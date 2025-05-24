package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/trades"
)

func (rest *rest) Trades(ctx context.Context, pair string) (*trades.GetResponse, error) {
	return rest.trades.GET(ctx, pair)
}
