package api

import (
	"io"
	"net/http"
)

type Client struct {
	*http.Client
}

// NewClient
func NewClient(c *http.Client) *Client {
	return &Client{c}
}

// (c *Client)Get ...
func (c *Client) GET(url string) (resp *http.Response, err error) {
	return c.Get(endpoint + url)
}

// (c *Client)POST new POST
func (c *Client) POST(url, bodyType string, body io.Reader) (resp *http.Response, err error) {
	return c.Post(endpoint+url, bodyType, body)
}
