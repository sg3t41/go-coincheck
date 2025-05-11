package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (c *rest) OpenOrders(ctx context.Context) (*output.Opens, error) {
	return c.opens.GET(ctx)
}
