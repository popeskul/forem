package forem

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) User(ctx context.Context) (UserResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, endpointUserMe, nil)
	if err != nil {
		return UserResponse{}, err
	}

	var user UserResponse
	resp, err := c.do(req, &user)
	if err != nil {
		return UserResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return UserResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return user, nil
}
