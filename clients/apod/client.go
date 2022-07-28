package apod

import (
	"errors"
	"net/http"
)

const baseURL = "https://api.nasa.gov/planetary/apod"

var ErrEmptyAPIKey = errors.New("api key cannot be empty")

type Option interface {
	apply(client *Client)
}

type OptionFn func(client *Client)

func (fn OptionFn) apply(client *Client) {
	fn(client)
}

func WithHTTPClient(httpClient *http.Client) OptionFn {
	return func(client *Client) {
		client.client = httpClient
	}
}

type Client struct {
	apiKey string

	client *http.Client
}

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	client := &Client{
		apiKey: apiKey,
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt.apply(client)
	}

	return client, nil
}
