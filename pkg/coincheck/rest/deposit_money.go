package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/depositmoney"
)

func (rest *rest) DepositMoneyHistory(ctx context.Context) (*depositmoney.GetResponse, error) {
	return rest.deposit_money.GET(ctx)
}

func (rest *rest) DepositMoneyFast(ctx context.Context, id int) (*depositmoney.PostFastResponse, error) {
	return rest.deposit_money.PostFast(ctx, id)
}