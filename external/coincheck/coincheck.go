package coincheck

import (
	"context"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/api/http/accounts"
	"github.com/sg3t41/go-coincheck/internal/api/http/accounts/balance"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchange/orders"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchange/orders/cancelstatus"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchange/orders/opens"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchange/orders/ordersrate"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchange/orders/transactions"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchange/orders/transactionspagination"
	"github.com/sg3t41/go-coincheck/internal/api/http/exchangestatus"
	"github.com/sg3t41/go-coincheck/internal/api/http/orderbooks"
	"github.com/sg3t41/go-coincheck/internal/api/http/referencerate"
	"github.com/sg3t41/go-coincheck/internal/api/http/ticker"
	"github.com/sg3t41/go-coincheck/internal/api/http/trades"
	"github.com/sg3t41/go-coincheck/internal/client"

	ws_trades "github.com/sg3t41/go-coincheck/internal/api/ws/trades"
)

type Coincheck interface {
	// rests
	Ticker(context.Context, input.GetTicker) (*output.GetTicker, error)
	Accounts(context.Context) (*output.Accounts, error)
	Balance(context.Context) (*output.Balance, error)
	ExchangeStatus(context.Context, input.ExchangeStatus) (*output.ExchangeStatus, error)
	ReferenceRate(context.Context, input.ReferenceRate) (*output.ReferenceRate, error)
	OrdersRate(context.Context, input.OrdersRate) (*output.OrdersRate, error)

	Trades(context.Context, input.GetTrades) (*output.GetTrades, error)
	OrderBooks(context.Context, input.GetOrderBooks) (*output.GetOrderBooks, error)
	GetOrder(context.Context, input.GetOrder) (*output.GetOrder, error)
	Transactions(context.Context) (*output.GetTransactions, error)
	TransactionsPagination(context.Context, input.TransactionsPagination) (*output.TransactionsPagination, error)
	OpenOrders(context.Context) (*output.Opens, error)

	CreateOrder(context.Context, input.CreateOrder) (*output.CreateOrder, error)
	CancelOrder(context.Context, input.CancelOrder) (*output.CancelOrder, error)
	CancelStatus(context.Context, input.CancelStatus) (*output.CancelStatus, error)

	// websockets
	WebSocketTrade(context context.Context, channel string, to chan<- string) error
}

type coincheck struct {
	// rests
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

	// websockets
	ws_trades ws_trades.WSTrades
}

func New(key, secret string) (Coincheck, error) {
	c, err := client.New(key, secret)
	if err != nil {
		return nil, err
	}

	return &coincheck{
		//rests
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

		// websockets
		ws_trades: ws_trades.New(c),
	}, nil
}
