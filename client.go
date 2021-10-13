package fetch

import (
	"net/http"
	"time"
)

// New client...
func DefaultClient() *Client {
	c := &Client{
		Name:    "default",
		UA:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
		Timeout: time.Minute,
	}
	c.init()
	return c
}

type Client struct {
	Name     string
	UA       string
	Timeout  time.Duration
	Insecure bool
	client   *http.Client
}
