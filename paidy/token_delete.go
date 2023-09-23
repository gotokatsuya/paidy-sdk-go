package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type TokenDeleteRequest struct {
	WalletID string `json:"wallet_id"`
	Reason   Reason `json:"reason"`
}

type TokenDeleteResponse Token

// トークンの削除
func (c Client) TokenDelete(ctx context.Context, id string, req *TokenDeleteRequest) (*TokenDeleteResponse, error) {
	path := fmt.Sprintf("/tokens/%s/delete", id)
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}
	resp := new(TokenDeleteResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
