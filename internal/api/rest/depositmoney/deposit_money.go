package depositmoney

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

type DepositMoney interface {
	// 入金履歴取得
	GET(ctx context.Context) (*GetResponse, error)
	// 高速入金
	PostFast(ctx context.Context, id int) (*PostFastResponse, error)
}

type depositMoney struct {
	client client.Client
}

func New(client client.Client) DepositMoney {
	return &depositMoney{client}
}

// DepositTransaction は入金取引を表す構造体
type DepositTransaction struct {
	ID        int       `json:"id"`         // 入金ID
	Amount    string    `json:"amount"`     // 入金量
	Currency  string    `json:"currency"`   // 通貨
	Address   string    `json:"address"`    // 入金アドレス
	TxHash    *string   `json:"txhash"`     // トランザクションハッシュ
	Status    string    `json:"status"`     // ステータス (pending, confirmed, etc.)
	CreatedAt time.Time `json:"created_at"` // 作成日時
	UpdatedAt time.Time `json:"updated_at"` // 更新日時
}

// GetResponse は入金履歴のレスポンス
type GetResponse struct {
	Success  bool                 `json:"success"`  // 成功フラグ
	Deposits []DepositTransaction `json:"deposits"` // 入金履歴
}

// PostFastResponse は高速入金のレスポンス
type PostFastResponse struct {
	Success bool   `json:"success"` // 成功フラグ
	ID      int    `json:"id"`      // 入金ID
	Status  string `json:"status"`  // ステータス
}

func (d depositMoney) GET(ctx context.Context) (*GetResponse, error) {
	req, err := d.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/deposit_money",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := d.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (d depositMoney) PostFast(ctx context.Context, id int) (*PostFastResponse, error) {
	bodyJSON, err := json.Marshal(map[string]string{
		"id": strconv.Itoa(id),
	})
	if err != nil {
		return nil, err
	}

	req, err := d.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodPost,
		Path:    "/api/deposit_money/" + strconv.Itoa(id) + "/fast",
		Body:    bytes.NewBuffer(bodyJSON),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res PostFastResponse
	if err := d.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}