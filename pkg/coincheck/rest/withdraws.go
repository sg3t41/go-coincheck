package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/withdraws"
)

func (rest *rest) Withdraws(ctx context.Context) (*withdraws.GetResponse, error) {
	return rest.withdrawals.GET(ctx)
}

func (rest *rest) CreateWithdraw(ctx context.Context, params withdraws.CreateWithdrawParams) (*withdraws.PostResponse, error) {
	return rest.withdrawals.POST(ctx, params)
}

func (rest *rest) CancelWithdraw(ctx context.Context, id int) (*withdraws.DeleteResponse, error) {
	return rest.withdrawals.DELETE(ctx, id)
}