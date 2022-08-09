package forem

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	endpointUserMe   = "api/users/me"
	endpointArticles = "api/articles"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "dev.to",
	Path:   "/",
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

type Client struct {
	httpClient *http.Client
	baseURL    url.URL
	ApiKey     string
	Timeout    time.Duration
}

func NewClient(client Client) (*Client, error) {
	if client.Timeout == 0 {
		client.Timeout = time.Second * 10
	}

	if client.ApiKey == "" {
		return nil, fmt.Errorf("api key is required")
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: client.Timeout,
			Transport: &LoggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
		ApiKey:  client.ApiKey,
		baseURL: baseURL,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	u := c.baseURL.ResolveReference(&url.URL{
		Path: path,
	})

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	req.WithContext(ctx)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("api-key", c.ApiKey)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}
