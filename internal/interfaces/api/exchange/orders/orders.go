// 参照: https://coincheck.com/ja/documents/exchange/api#Ticker
package orders

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sg3t41/go-coincheck/external/dto/input"
	"github.com/sg3t41/go-coincheck/external/dto/output"
	"github.com/sg3t41/go-coincheck/internal/infrastructure/client"
)

type Orders interface {
	// 新規注文
	POST(context.Context, input.CreateOrder) (*output.CreateOrder, error)
	DELETE(context.Context, input.CancelOrder) (*output.CancelOrder, error)
	GET(context.Context, input.GetOrder) (*output.GetOrder, error)
}

type orders struct {
	client client.Client
}

func New(client client.Client) Orders {
	return &orders{
		client,
	}
}

func (o orders) GET(ctx context.Context, in input.GetOrder) (*output.GetOrder, error) {
	req, err := o.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/exchange/orders/" + strconv.Itoa(in.ID),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.GetOrder
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil

}

func (o orders) POST(ctx context.Context, in input.CreateOrder) (*output.CreateOrder, error) {
	bodyJSON, err := json.Marshal(
		map[string]string{
			"order_type": in.OrderType,
			"amount":     strconv.FormatFloat(in.Amount, 'f', -1, 64),
			"rate":       strconv.FormatFloat(in.Rate, 'f', -1, 64),
			"pair":       in.Pair,
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := o.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodPost,
		Path:    "/api/exchange/orders",
		Body:    bytes.NewBuffer(bodyJSON),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.CreateOrder
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (o orders) DELETE(ctx context.Context, in input.CancelOrder) (*output.CancelOrder, error) {
	req, err := o.client.CreateRequest(ctx, client.RequestInput{
		Method:  http.MethodDelete,
		Path:    "/api/exchange/orders/" + strconv.Itoa(in.ID),
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res output.CancelOrder
	if err := o.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
