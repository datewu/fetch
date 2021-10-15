package fetch

import (
	"context"
	"net/http"
	"time"
)

// New client...
func DefaultClient(ctx context.Context) *Client {
	c := &Client{
		Name:    "default",
		UA:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
		Timeout: time.Minute,
		Retry:   3,
		ctx:     ctx,
	}
	c.init()
	return c
}

type Client struct {
	Name     string
	UA       string
	Timeout  time.Duration
	Retry    int
	Insecure bool
	client   *http.Client
	ctx      context.Context
}

func (c *Client) retry(f func() error) error {
	if c.Retry <= 0 || c.ctx == nil {
		return f()
	}
	var err error
	for i := 0; i < c.Retry; i++ {
		err = f()
		if err == nil {
			return nil
		}
		select {
		case <-c.ctx.Done():
			return c.ctx.Err()
		default:
			time.Sleep(2 << (i + 1) * time.Second)
		}
	}
	return err
}
