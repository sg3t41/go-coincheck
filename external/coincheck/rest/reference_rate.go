package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) ReferenceRate(ctx context.Context, in input.ReferenceRate) (*output.ReferenceRate, error) {
	return c.reference_rate.GET(ctx, in)
}
