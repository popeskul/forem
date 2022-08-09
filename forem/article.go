package forem

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	UnexpectedStatusCode = errors.New("unexpected status code")
)

func (c *Client) CreateArticle(ctx context.Context, input Article) (ArticleResponse, error) {
	if err := input.validate(); err != nil {
		return ArticleResponse{}, err
	}

	body, err := json.Marshal(ArticleInput{Article: input})

	req, err := c.newRequest(ctx, http.MethodPost, endpointArticles, body)
	if err != nil {
		return ArticleResponse{}, err
	}

	var articleResponse ArticleResponse
	resp, err := c.do(req, &articleResponse)
	if err != nil {
		return ArticleResponse{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return ArticleResponse{}, UnexpectedStatusCode
	}

	return articleResponse, nil
}
