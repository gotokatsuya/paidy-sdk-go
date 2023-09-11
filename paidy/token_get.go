package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type TokenGetResponse Token

// 特定のtokenオブジェクトを取得
func (c Client) TokenGet(ctx context.Context, id string) (*TokenGetResponse, error) {
	path := fmt.Sprintf("/tokens/%s", id)
	httpReq, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp := new(TokenGetResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
