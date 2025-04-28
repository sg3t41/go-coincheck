package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

// GetOpenOrders は、オープンオーダーを取得します
func (c *coincheck) OpenOrders(ctx context.Context) (*output.Opens, error) {
	return c.opens.GET(ctx)
}
