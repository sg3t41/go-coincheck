package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sg3t41/go-coincheck/external/e"
)

type HTTPClient interface {
	CreateRequest(ctx context.Context, input RequestInput) (*http.Request, error)
	Do(req *http.Request, output any) error
}

type httpClient struct {
	httpClient  *http.Client
	baseURL     *url.URL
	credentials *credentials
}

func NewHTTPClient(key, secret string) (HTTPClient, error) {
	baseURL, err := url.Parse("https://coincheck.com")
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	return &httpClient{
		httpClient:  http.DefaultClient,
		credentials: &credentials{key, secret},
		baseURL:     baseURL,
	}, nil
}

type RequestInput struct {
	Method     string            // HTTPメソッド (例: GET, POST)
	Path       string            // APIパス (例: /api/orders)
	Body       io.Reader         // リクエストボディ。不要な場合はnilを設定します。
	QueryParam map[string]string // クエリパラメータ (例: {"pair": "btc_jpy"}) 不要な場合はnilを設定します。
	Private    bool              // trueの場合、プライベートAPIです。
}

func (c *httpClient) CreateRequest(ctx context.Context, input RequestInput) (*http.Request, error) {
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

func (c *httpClient) Do(req *http.Request, output any) error {
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

func (c *httpClient) hasCredentials() bool {
	return c.credentials != nil
}

func (c *httpClient) setAuthHeaders(req *http.Request, body string) error {
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
