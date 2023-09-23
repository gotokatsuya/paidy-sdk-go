package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type PaymentRefundRequest struct {
	CaptureID string `json:"capture_id,omitempty"`
	Amount    int    `json:"amount,omitempty"`
}

type PaymentRefundResponse Payment

// 決済のRefund
func (c Client) PaymentRefund(ctx context.Context, id string, req *PaymentRefundRequest) (*PaymentRefundResponse, error) {
	path := fmt.Sprintf("/payments/%s/refunds", id)
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}
	resp := new(PaymentRefundResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
