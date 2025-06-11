package rest

import (
	"context"

	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts"
	"github.com/sg3t41/go-coincheck/internal/api/rest/accounts/balance"
	"github.com/sg3t41/go-coincheck/internal/api/rest/bankaccounts"
	"github.com/sg3t41/go-coincheck/internal/api/rest/depositmoney"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/cancelstatus"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/opens"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/ordersrate"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactions"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactionspagination"
	"github.com/sg3t41/go-coincheck/internal/api/rest/exchangestatus"
	"github.com/sg3t41/go-coincheck/internal/api/rest/orderbooks"
	"github.com/sg3t41/go-coincheck/internal/api/rest/referencerate"
	"github.com/sg3t41/go-coincheck/internal/api/rest/sendmoney"
	"github.com/sg3t41/go-coincheck/internal/api/rest/ticker"
	"github.com/sg3t41/go-coincheck/internal/api/rest/trades"
	"github.com/sg3t41/go-coincheck/internal/api/rest/withdraws"
	"github.com/sg3t41/go-coincheck/internal/client"
)

type REST interface {
	// Public APIs
	Ticker(ctx context.Context, pair string) (*ticker.GetResponse, error)
	Trades(ctx context.Context, pair string) (*trades.GetResponse, error)
	OrderBooks(ctx context.Context, pair string) (*orderbooks.GetResponse, error)
	OrdersRate(ctx context.Context, pair, orderType string, price, amount float64) (*ordersrate.GetResponse, error)
	ReferenceRate(ctx context.Context, pair string) (*referencerate.GetResponse, error)
	ExchangeStatus(ctx context.Context, pair string) (*exchangestatus.GetResponse, error)

	// Account APIs
	Accounts(context.Context) (*accounts.GetResponse, error)
	Balance(context.Context) (*balance.GetResponse, error)

	// Trading APIs
	Transactions(context.Context) (*transactions.GetResponse, error)
	TransactionsPagination(ctx context.Context, limit int, order string, startingAfter, endingBefore *int) (*transactionspagination.GetResponse, error)
	OpenOrders(context.Context) (*opens.GetResponse, error)
	GetOrder(ctx context.Context, id int) (*orders.GetResponse, error)
	CreateOrder(ctx context.Context, pair, orderType string, rate, amount float64) (*orders.PostResponse, error)
	CancelOrder(ctx context.Context, id int) (*orders.DeleteResponse, error)
	CancelStatus(ctx context.Context, id int) (*cancelstatus.GetResponse, error)

	// Transfer APIs
	SendMoney(ctx context.Context, params sendmoney.SendMoneyParams) (*sendmoney.PostResponse, error)
	SendMoneyHistory(ctx context.Context) (*sendmoney.GetResponse, error)
	DepositMoneyHistory(ctx context.Context) (*depositmoney.GetResponse, error)
	DepositMoneyFast(ctx context.Context, id int) (*depositmoney.PostFastResponse, error)

	// Bank Account APIs
	BankAccounts(ctx context.Context) (*bankaccounts.GetResponse, error)
	CreateBankAccount(ctx context.Context, params bankaccounts.CreateBankAccountParams) (*bankaccounts.PostResponse, error)
	DeleteBankAccount(ctx context.Context, id int) (*bankaccounts.DeleteResponse, error)

	// Withdraw APIs
	Withdraws(ctx context.Context) (*withdraws.GetResponse, error)
	CreateWithdraw(ctx context.Context, params withdraws.CreateWithdrawParams) (*withdraws.PostResponse, error)
	CancelWithdraw(ctx context.Context, id int) (*withdraws.DeleteResponse, error)
}

type rest struct {
	// Account related
	accounts accounts.Accounts
	balance  balance.Balance

	// Trading related
	cancel_status           cancelstatus.CancelStatus
	orders                  orders.Orders
	exchange_status         exchangestatus.ExchangeStatus
	orders_rate             ordersrate.OrdersRate
	opens                   opens.Opens
	trades                  trades.Trades
	transactions            transactions.Transactions
	transactions_pagination transactionspagination.TransactionsPagination
	order_books             orderbooks.OrderBooks
	reference_rate          referencerate.ReferenceRate
	ticker                  ticker.Ticker

	// Transfer related
	send_money    sendmoney.SendMoney
	deposit_money depositmoney.DepositMoney

	// Bank account related
	bank_accounts bankaccounts.BankAccounts

	// Withdraw related
	withdrawals withdraws.Withdraws
}

func New(key, secret string) (REST, error) {
	c, err := client.New(
		client.WithREST(key, secret, "https://coincheck.com"), // RESTクライアントのみを初期化
	)
	if err != nil {
		return nil, err
	}

	return &rest{
		// Account related
		accounts: accounts.New(c),
		balance:  balance.New(c),

		// Trading related
		cancel_status:           cancelstatus.New(c),
		orders:                  orders.New(c),
		exchange_status:         exchangestatus.New(c),
		orders_rate:             ordersrate.New(c),
		opens:                   opens.New(c),
		trades:                  trades.New(c),
		transactions:            transactions.New(c),
		transactions_pagination: transactionspagination.New(c),
		order_books:             orderbooks.New(c),
		reference_rate:          referencerate.New(c),
		ticker:                  ticker.New(c),

		// Transfer related
		send_money:    sendmoney.New(c),
		deposit_money: depositmoney.New(c),

		// Bank account related
		bank_accounts: bankaccounts.New(c),

		// Withdraw related
		withdrawals: withdraws.New(c),
	}, nil
}
