package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type PaymentGetResponse Payment

// 決済データの取得
func (c Client) PaymentGet(ctx context.Context, id string) (*PaymentGetResponse, error) {
	path := fmt.Sprintf("/payments/%s", id)
	httpReq, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PaymentGetResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
