package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

// RequestHeaderParam はリクエストヘッダーのパラメータを表します。
type RequestHeaderParam struct {
	AccessKey       string
	AccessNonce     string
	AccessSignature string
}

// Credentials はAPIキーとシークレットを保持する構造体です。
type credentials struct {
	key    string
	secret string
}

// GenerateRequestHeaders はリクエストに必要なヘッダーを生成します。
func (c *credentials) GenerateRequestHeaders(requestURL *url.URL, body string) (*RequestHeaderParam, error) {
	nonce := time.Now().UnixNano() / int64(time.Millisecond) // nonceをミリ秒単位で生成
	message := fmt.Sprintf("%d%s%s", nonce, requestURL.String(), body)

	h := hmac.New(sha256.New, []byte(c.secret))
	if _, err := h.Write([]byte(message)); err != nil {
		return nil, fmt.Errorf("failed to write HMAC: %w", err)
	}
	signature := hex.EncodeToString(h.Sum(nil))

	return &RequestHeaderParam{
		AccessKey:       c.key,
		AccessNonce:     fmt.Sprintf("%d", nonce),
		AccessSignature: signature,
	}, nil
}
