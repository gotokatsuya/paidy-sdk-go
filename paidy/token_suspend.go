package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type TokenSuspendRequest struct {
	WalletID string `json:"wallet_id"`
	Reason   Reason `json:"reason"`
}

type Reason struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type TokenSuspendResponse Token

// トークンの一時的な無効化
func (c Client) TokenSuspend(ctx context.Context, id string, req *TokenSuspendRequest) (*TokenSuspendResponse, error) {
	path := fmt.Sprintf("/tokens/%s/suspend", id)
	httpReq, err := c.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	resp := new(TokenSuspendResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
