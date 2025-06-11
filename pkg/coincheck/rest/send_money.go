package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/sendmoney"
)

func (rest *rest) SendMoney(ctx context.Context, params sendmoney.SendMoneyParams) (*sendmoney.PostResponse, error) {
	return rest.send_money.POST(ctx, params)
}

func (rest *rest) SendMoneyHistory(ctx context.Context) (*sendmoney.GetResponse, error) {
	return rest.send_money.GET(ctx)
}