package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/output"
)

func (coincheck *coincheck) Accounts(ctx context.Context) (*output.Accounts, error) {
	return coincheck.accounts.Get(ctx)
}
