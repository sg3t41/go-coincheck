package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/referencerate"
)

func (c *rest) ReferenceRate(ctx context.Context, pair string) (*referencerate.GetResponse, error) {
	return c.reference_rate.GET(ctx, pair)
}
