package withdraws

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/sg3t41/go-coincheck/internal/client"
	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

type Withdraws interface {
	// 出金履歴取得
	GET(ctx context.Context) (*GetResponse, error)
	// 出金申請
	POST(ctx context.Context, params CreateWithdrawParams) (*PostResponse, error)
	// 出金申請キャンセル
	DELETE(ctx context.Context, id int) (*DeleteResponse, error)
}

type withdraws struct {
	client client.Client
}

func New(client client.Client) Withdraws {
	return &withdraws{client}
}

// CreateWithdrawParams は出金申請時のパラメータ
type CreateWithdrawParams struct {
	BankAccountID int    `json:"bank_account_id"` // 銀行口座ID (必須)
	Amount        string `json:"amount"`          // 出金金額 (必須)
	Currency      string `json:"currency"`        // 通貨 (jpy固定)
	IsFast        bool   `json:"is_fast"`         // 高速出金フラグ (オプション)
}

// WithdrawTransaction は出金取引を表す構造体
type WithdrawTransaction struct {
	ID            int       `json:"id"`              // 出金ID
	Status        string    `json:"status"`          // ステータス (pending, processing, completed, etc.)
	Amount        string    `json:"amount"`          // 出金金額
	Currency      string    `json:"currency"`        // 通貨
	Fee           string    `json:"fee"`             // 手数料
	IsFast        bool      `json:"is_fast"`         // 高速出金フラグ
	BankAccountID int       `json:"bank_account_id"` // 銀行口座ID
	CreatedAt     time.Time `json:"created_at"`      // 作成日時
	UpdatedAt     time.Time `json:"updated_at"`      // 更新日時
}

// GetResponse は出金履歴のレスポンス
type GetResponse struct {
	Success   bool                  `json:"success"`   // 成功フラグ
	Withdraws []WithdrawTransaction `json:"withdraws"` // 出金履歴
}

// PostResponse は出金申請のレスポンス
type PostResponse struct {
	Success       bool   `json:"success"`         // 成功フラグ
	ID            int    `json:"id"`              // 出金ID
	Amount        string `json:"amount"`          // 出金金額
	Currency      string `json:"currency"`        // 通貨
	Fee           string `json:"fee"`             // 手数料
	IsFast        bool   `json:"is_fast"`         // 高速出金フラグ
	BankAccountID int    `json:"bank_account_id"` // 銀行口座ID
	Status        string `json:"status"`          // ステータス
}

// DeleteResponse は出金申請キャンセルのレスポンス
type DeleteResponse struct {
	Success bool `json:"success"` // 成功フラグ
	ID      int  `json:"id"`      // キャンセルされた出金ID
}

func (w withdraws) GET(ctx context.Context) (*GetResponse, error) {
	req, err := w.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/withdraws",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := w.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (w withdraws) POST(ctx context.Context, params CreateWithdrawParams) (*PostResponse, error) {
	// デフォルト通貨設定
	if params.Currency == "" {
		params.Currency = "jpy"
	}

	body := map[string]interface{}{
		"bank_account_id": params.BankAccountID,
		"amount":          params.Amount,
		"currency":        params.Currency,
	}

	if params.IsFast {
		body["is_fast"] = "true"
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := w.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodPost,
		Path:    "/api/withdraws",
		Body:    bytes.NewBuffer(bodyJSON),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res PostResponse
	if err := w.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (w withdraws) DELETE(ctx context.Context, id int) (*DeleteResponse, error) {
	req, err := w.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodDelete,
		Path:    "/api/withdraws/" + strconv.Itoa(id),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	if err := w.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}