package repository

import (
	"context"
	"net/http"

	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/domain"
	"github.com/sg3t41/go-coincheck/internal/infrastructure/client"
)

type balance struct {
	client client.Client
}

func NewBalance(client client.Client) domain.BalanceRepository {
	return &balance{
		client: client,
	}
}

func (balance balance) Get(ctx context.Context) (*output.Balance, error) {
	req, err := balance.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts/balance",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.Balance
	if err := balance.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
