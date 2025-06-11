package sendmoney

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

type SendMoney interface {
	// 送金実行
	POST(ctx context.Context, params SendMoneyParams) (*PostResponse, error)
	// 送金履歴取得
	GET(ctx context.Context) (*GetResponse, error)
}

type sendMoney struct {
	client client.Client
}

func New(client client.Client) SendMoney {
	return &sendMoney{client}
}

// SendMoneyParams は送金時のパラメータ
type SendMoneyParams struct {
	Address  string  `json:"address"`            // 送金先アドレス (必須)
	Amount   float64 `json:"amount"`             // 送金量 (必須)
	Currency string  `json:"currency,omitempty"` // 通貨 (デフォルト: btc)
}

// SendMoneyTransaction は送金取引を表す構造体
type SendMoneyTransaction struct {
	ID          int       `json:"id"`           // 送金ID
	Address     string    `json:"address"`      // 送金先アドレス
	Amount      string    `json:"amount"`       // 送金量
	Fee         string    `json:"fee"`          // 手数料
	Currency    string    `json:"currency"`     // 通貨
	TxHash      *string   `json:"txhash"`       // トランザクションハッシュ
	Status      string    `json:"status"`       // ステータス
	CreatedAt   time.Time `json:"created_at"`   // 作成日時
	UpdatedAt   time.Time `json:"updated_at"`   // 更新日時
}

// PostResponse は送金実行のレスポンス
type PostResponse struct {
	Success bool   `json:"success"` // 成功フラグ
	ID      int    `json:"id"`      // 送金ID
	Address string `json:"address"` // 送金先アドレス
	Amount  string `json:"amount"`  // 送金量
	Fee     string `json:"fee"`     // 手数料
}

// GetResponse は送金履歴のレスポンス
type GetResponse struct {
	Success      bool                   `json:"success"`      // 成功フラグ
	SendMoneyies []SendMoneyTransaction `json:"send_moneys"`  // 送金履歴
}

func (s sendMoney) POST(ctx context.Context, params SendMoneyParams) (*PostResponse, error) {
	// デフォルト通貨設定
	if params.Currency == "" {
		params.Currency = "btc"
	}

	bodyJSON, err := json.Marshal(map[string]string{
		"address":  params.Address,
		"amount":   strconv.FormatFloat(params.Amount, 'f', -1, 64),
		"currency": params.Currency,
	})
	if err != nil {
		return nil, err
	}

	req, err := s.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodPost,
		Path:    "/api/send_money",
		Body:    bytes.NewBuffer(bodyJSON),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res PostResponse
	if err := s.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s sendMoney) GET(ctx context.Context) (*GetResponse, error) {
	req, err := s.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/send_money",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := s.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}