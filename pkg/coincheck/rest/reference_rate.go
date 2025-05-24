package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/referencerate"
)

func (rest *rest) ReferenceRate(ctx context.Context, pair string) (*referencerate.GetResponse, error) {
	return rest.reference_rate.GET(ctx, pair)
}
