package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type PaymentCloseRequest struct{}

type PaymentCloseResponse Payment

// 決済のClose
func (c Client) PaymentClose(ctx context.Context, id string) (*PaymentCloseResponse, error) {
	path := fmt.Sprintf("/payments/%s/close", id)
	httpReq, err := c.NewRequest(http.MethodPost, path, &PaymentCloseRequest{})
	if err != nil {
		return nil, err
	}
	resp := new(PaymentCloseResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
