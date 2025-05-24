package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/ticker"
)

func (rest *rest) Ticker(ctx context.Context, pair string) (*ticker.GetResponse, error) {
	return rest.ticker.GET(ctx, pair)
}
