package bankaccounts

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

type BankAccounts interface {
	// 銀行口座一覧取得
	GET(ctx context.Context) (*GetResponse, error)
	// 銀行口座登録
	POST(ctx context.Context, params CreateBankAccountParams) (*PostResponse, error)
	// 銀行口座削除
	DELETE(ctx context.Context, id int) (*DeleteResponse, error)
}

type bankAccounts struct {
	client client.Client
}

func New(client client.Client) BankAccounts {
	return &bankAccounts{client}
}

// CreateBankAccountParams は銀行口座登録時のパラメータ
type CreateBankAccountParams struct {
	BankName       string `json:"bank_name"`       // 銀行名 (必須)
	BranchName     string `json:"branch_name"`     // 支店名 (必須)
	BankType       string `json:"bank_type"`       // 口座種別 (futsu: 普通, touza: 当座)
	Number         string `json:"number"`          // 口座番号 (必須)
	Name           string `json:"name"`            // 口座名義 (必須)
}

// BankAccount は銀行口座情報を表す構造体
type BankAccount struct {
	ID         int       `json:"id"`          // 口座ID
	BankName   string    `json:"bank_name"`   // 銀行名
	BranchName string    `json:"branch_name"` // 支店名
	BankType   string    `json:"bank_type"`   // 口座種別
	Number     string    `json:"number"`      // 口座番号
	Name       string    `json:"name"`        // 口座名義
	CreatedAt  time.Time `json:"created_at"`  // 作成日時
	UpdatedAt  time.Time `json:"updated_at"`  // 更新日時
}

// GetResponse は銀行口座一覧のレスポンス
type GetResponse struct {
	Success      bool          `json:"success"`       // 成功フラグ
	BankAccounts []BankAccount `json:"bank_accounts"` // 銀行口座一覧
}

// PostResponse は銀行口座登録のレスポンス
type PostResponse struct {
	Success     bool   `json:"success"`      // 成功フラグ
	ID          int    `json:"id"`           // 口座ID
	BankName    string `json:"bank_name"`    // 銀行名
	BranchName  string `json:"branch_name"`  // 支店名
	BankType    string `json:"bank_type"`    // 口座種別
	Number      string `json:"number"`       // 口座番号
	Name        string `json:"name"`         // 口座名義
}

// DeleteResponse は銀行口座削除のレスポンス
type DeleteResponse struct {
	Success bool `json:"success"` // 成功フラグ
	ID      int  `json:"id"`      // 削除された口座ID
}

func (b bankAccounts) GET(ctx context.Context) (*GetResponse, error) {
	req, err := b.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/bank_accounts",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	if err := b.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (b bankAccounts) POST(ctx context.Context, params CreateBankAccountParams) (*PostResponse, error) {
	bodyJSON, err := json.Marshal(map[string]string{
		"bank_name":   params.BankName,
		"branch_name": params.BranchName,
		"bank_type":   params.BankType,
		"number":      params.Number,
		"name":        params.Name,
	})
	if err != nil {
		return nil, err
	}

	req, err := b.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodPost,
		Path:    "/api/bank_accounts",
		Body:    bytes.NewBuffer(bodyJSON),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res PostResponse
	if err := b.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (b bankAccounts) DELETE(ctx context.Context, id int) (*DeleteResponse, error) {
	req, err := b.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodDelete,
		Path:    "/api/bank_accounts/" + strconv.Itoa(id),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	if err := b.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}