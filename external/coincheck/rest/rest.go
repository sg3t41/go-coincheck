package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts"
	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts/balance"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/cancelstatus"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/opens"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/ordersrate"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactions"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactionspagination"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchangestatus"
	"github.com/sg3t41/go-coincheck/internal/api/rest/orderbooks"
	"github.com/sg3t41/go-coincheck/internal/api/rest/referencerate"
	"github.com/sg3t41/go-coincheck/internal/api/rest/ticker"
	"github.com/sg3t41/go-coincheck/internal/api/rest/trades"
	"github.com/sg3t41/go-coincheck/internal/client"
)

type REST interface {
	// rests
	Ticker(ctx context.Context, pair string) (*ticker.GetResponse, error)
	Accounts(context.Context) (*accounts.GetResponse, error)
	Balance(context.Context) (*balance.GetReponse, error)
	ExchangeStatus(ctx context.Context, pair string) (*exchangestatus.GetReponse, error)
	ReferenceRate(ctx context.Context, pair string) (*referencerate.GetResponse, error)
	OrdersRate(ctx context.Context, pair, orderType string, price, amount float64) (*ordersrate.GetResponse, error)

	Trades(ctx context.Context, pair string) (*trades.GetResponse, error)
	OrderBooks(ctx context.Context, pair string) (*orderbooks.GetResponse, error)
	Transactions(context.Context) (*transactions.GetReponse, error)
	TransactionsPagination(ctx context.Context, limit int, order string, startingAfter, endingBefore *int) (*transactionspagination.GetResponse, error)
	OpenOrders(context.Context) (*opens.GetResponse, error)

	GetOrder(ctx context.Context, id int) (*orders.GetResponse, error)
	CreateOrder(ctx context.Context, pair, orderType string, rate, amount float64) (*orders.PostResponse, error)
	CancelOrder(ctx context.Context, id int) (*orders.DeleteResponse, error)

	CancelStatus(ctx context.Context, id int) (*cancelstatus.GetResponse, error)
}

type rest struct {
	accounts                accounts.Accounts
	cancel_status           cancelstatus.CancelStatus
	orders                  orders.Orders
	exchange_status         exchangestatus.ExchangeStatus
	orders_rate             ordersrate.OrdersRate
	opens                   opens.Opens
	balance                 balance.Balance
	trades                  trades.Trades
	transactions            transactions.Transactions
	transactions_pagination transactionspagination.TransactionsPagination
	order_books             orderbooks.OrderBooks
	reference_rate          referencerate.ReferenceRate
	ticker                  ticker.Ticker
}

func New(key, secret string) (REST, error) {
	c, err := client.New(key, secret)
	if err != nil {
		return nil, err
	}

	return &rest{
		accounts:                accounts.New(c),
		balance:                 balance.New(c),
		cancel_status:           cancelstatus.New(c),
		orders:                  orders.New(c),
		exchange_status:         exchangestatus.New(c),
		orders_rate:             ordersrate.New(c),
		opens:                   opens.New(c),
		trades:                  trades.New(c),
		transactions:            transactions.New(c),
		transactions_pagination: transactionspagination.New(c),
		reference_rate:          referencerate.New(c),
		ticker:                  ticker.New(c),
	}, nil
}
