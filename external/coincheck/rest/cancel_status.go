package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) CancelStatus(ctx context.Context, in input.CancelStatus) (*output.CancelStatus, error) {
	return c.cancel_status.GET(ctx, in)
}
