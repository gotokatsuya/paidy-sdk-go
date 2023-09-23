package paidy

import (
	"context"
	"fmt"
	"net/http"
)

type TokenResumeRequest struct {
	WalletID string `json:"wallet_id"`
	Reason   Reason `json:"reason"`
}

type TokenResumeResponse Token

// トークンの回復
func (c Client) TokenResume(ctx context.Context, id string, req *TokenResumeRequest) (*TokenResumeResponse, error) {
	path := fmt.Sprintf("/tokens/%s/resume", id)
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}
	resp := new(TokenResumeResponse)
	_, err = c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
