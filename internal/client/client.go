package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sg3t41/go-coincheck/internal/e"
)

const (
	// BaseURLはcoincheck APIのベースURLです。
	BaseURL = "https://coincheck.com"
)

// Clientはcoincheckクライアントを表します。
type client struct {
	// clientはCoincheck APIと通信するために使用されるHTTPクライアントです。
	httpClient *http.Client

	// websocketClient
	webSocketClient WebSocketClient

	// baseURLはcoincheck APIのベースURLです。
	baseURL *url.URL
	// credentialsはcoincheck APIに認証するために使用される資格情報です。
	credentials *credentials
}

type Client interface {
	CreateRequest(ctx context.Context, input RequestInput) (*http.Request, error)

	Do(req *http.Request, output any) error
	Connect(context.Context) error

	Subscribe(context.Context, string) chan any
}

func (c *client) setBaseURL(url *url.URL) {
	c.baseURL = url
}

func (c *client) setHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *client) setCredentials(credentials *credentials) {
	c.credentials = credentials
}

// NewClientは新しいcoincheckクライアントを返します。
func New(opts ...Option) (Client, error) {
	c := &client{}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.baseURL == nil {
		baseURL, err := url.Parse(BaseURL)
		if err != nil {
			return nil, e.WithPrefixError(err)
		}
		c.baseURL = baseURL
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c, nil
}

func (c *client) Connect(context.Context) error {
	return nil
}

func (c *client) Subscribe(context.Context, string) chan any {
	return nil
}

// hasCredentialsはクライアントに資格情報がある場合にtrueを返します。
func (c *client) hasCredentials() bool {
	return c.credentials != nil
}

func (c *client) setWebSocketClient(wsc WebSocketClient) {
	c.webSocketClient = wsc
}

// setAuthHeadersはリクエストに認証ヘッダーを設定します。
// プライベートAPIを使用する場合は、coincheckのウェブサイトからAPIキーとAPIシークレットを取得する必要があります。
func (c *client) setAuthHeaders(req *http.Request, body string) error {
	if !c.hasCredentials() {
		return e.ErrNoCredentials
	}

	headers, err := c.credentials.GenerateRequestHeaders(req.URL, body)
	if err != nil {
		return e.WithPrefixError(err)
	}
	req.Header.Set("ACCESS-KEY", headers.AccessKey)
	req.Header.Set("ACCESS-NONCE", headers.AccessNonce)
	req.Header.Set("ACCESS-SIGNATURE", headers.AccessSignature)
	return nil
}

// RequestInputはcreateRequestの入力パラメータを表します。
type RequestInput struct {
	Method     string            // HTTPメソッド (例: GET, POST)
	Path       string            // APIパス (例: /api/orders)
	Body       io.Reader         // リクエストボディ。不要な場合はnilを設定します。
	QueryParam map[string]string // クエリパラメータ (例: {"pair": "btc_jpy"}) 不要な場合はnilを設定します。
	Private    bool              // trueの場合、プライベートAPIです。
}

// createRequestは新しいHTTPリクエストを作成します。
func (c *client) CreateRequest(ctx context.Context, input RequestInput) (*http.Request, error) {
	u, err := url.JoinPath(c.baseURL.String(), input.Path)
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	endpoint, err := url.Parse(u)
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	if input.QueryParam != nil {
		q := endpoint.Query()
		for k, v := range input.QueryParam {
			q.Set(k, v)
		}
		endpoint.RawQuery = q.Encode()
	}

	var bodyStr string
	if input.Body != nil {
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, input.Body)
		if err != nil {
			return nil, e.WithPrefixError(err)
		}
		bodyStr = buf.String()
	}

	req, err := http.NewRequestWithContext(ctx, input.Method, endpoint.String(), strings.NewReader(bodyStr))
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	if input.Private {
		if err := c.setAuthHeaders(req, bodyStr); err != nil {
			return nil, err
		}
	}

	return req, nil
}

// DoはHTTPリクエストを送信し、HTTPレスポンスを返します。
func (c *client) Do(req *http.Request, output any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return e.WithPrefixError(err)
	}
	defer resp.Body.Close() //nolint: errcheck // エラーを無視

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return e.WithPrefixError(err)
	}

	if resp.StatusCode != http.StatusOK {
		return e.WithPrefixError(fmt.Errorf("予期しないステータスコード=%d, レスポンスボディ=%s", resp.StatusCode, string(bodyBytes)))
	}

	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(output); err != nil {
		return e.WithPrefixError(err)
	}
	return nil
}
