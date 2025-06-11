package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/bankaccounts"
)

func (rest *rest) BankAccounts(ctx context.Context) (*bankaccounts.GetResponse, error) {
	return rest.bank_accounts.GET(ctx)
}

func (rest *rest) CreateBankAccount(ctx context.Context, params bankaccounts.CreateBankAccountParams) (*bankaccounts.PostResponse, error) {
	return rest.bank_accounts.POST(ctx, params)
}

func (rest *rest) DeleteBankAccount(ctx context.Context, id int) (*bankaccounts.DeleteResponse, error) {
	return rest.bank_accounts.DELETE(ctx, id)
}