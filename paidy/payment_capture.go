package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type PaymentCaptureRequest struct{}

type PaymentCaptureResponse Payment

// 決済のCapture
func (c Client) PaymentCapture(ctx context.Context, id string) (*PaymentCaptureResponse, error) {
	path := fmt.Sprintf("/payments/%s/captures", id)
	httpReq, err := c.NewRequest(http.MethodPost, path, &PaymentCaptureRequest{})
	if err != nil {
		return nil, err
	}
	resp := new(PaymentCaptureResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
