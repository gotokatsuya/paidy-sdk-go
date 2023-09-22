package paidy

import (
	"context"
	"net/http"
)

type PaymentCreateRequest struct {
	TokenID         string          `json:"token_id"`
	Amount          int             `json:"amount"`
	Currency        string          `json:"currency"`
	Description     string          `json:"description,omitempty"`
	StoreName       string          `json:"store_name,omitempty"`
	BuyerData       BuyerData       `json:"buyer_data"`
	Order           Order           `json:"order"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type BuyerData struct {
	UserID          string `json:"user_id,omitempty"`
	Age             int    `json:"age"`
	OrderCount      int    `json:"order_count"`
	Ltv             int    `json:"ltv"`
	LastOrderAmount int    `json:"last_order_amount"`
	LastOrderAt     int    `json:"last_order_at"`
}

type Order struct {
	Items    []Item `json:"items"`
	Tax      int    `json:"tax,omitempty"`
	Shipping int    `json:"shipping,omitempty"`
	OrderRef string `json:"order_ref,omitempty"`
}

type Item struct {
	Quantity    int    `json:"quantity"`
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	UnitPrice   int    `json:"unit_price"`
}

type ShippingAddress struct {
	Line1 string `json:"line1,omitempty"`
	Line2 string `json:"line2,omitempty"`
	State string `json:"state,omitempty"`
	City  string `json:"city,omitempty"`
	Zip   string `json:"zip"`
}

type PaymentCreateResponse Payment

// 定期購入の決済を作成
func (c Client) PaymentCreate(ctx context.Context, req *PaymentCreateRequest) (*PaymentCreateResponse, error) {
	path := "/payments"
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}
	resp := new(PaymentCreateResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
