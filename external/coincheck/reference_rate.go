package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetRate は、指定されたペアの基準レートを取得します
func (c *coincheck) ReferenceRate(ctx context.Context, in input.ReferenceRate) (*output.ReferenceRate, error) {
	return c.reference_rate.GET(ctx, in)
}
