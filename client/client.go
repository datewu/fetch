package client

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/datewu/fetch/pattern"
	"github.com/datewu/fetch/request"
)

const (
	defaultUA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36"
)

// New client...
func DefaultClient(ctx context.Context) *Client {
	c := &Client{
		Name:    "spoof",
		UA:      defaultUA,
		Timeout: time.Minute,
		Retry:   3,
		ctx:     ctx,
	}
	c.setDefaultCli()
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

func (c *Client) SetHTTPClient(cli *http.Client) {
	if cli == nil {
		c.setDefaultCli()
	} else {
		c.client = cli
	}
}

func (c *Client) setDefaultCli() {
	cli := &http.Client{
		Timeout: c.Timeout,
	}
	if c.Insecure {
		cli.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	c.client = cli
}

func (c *Client) retry(f func() error) error {
	if c.ctx == nil {
		return f()
	}
	return pattern.Retry(c.ctx, f, c.Retry)
}

func (c *Client) newRequsetParams(metod, url string, r io.Reader, modify request.ModifyFunc) *request.ReqParams {
	params := request.NewReqParams(c.ctx)
	params.Method = metod
	params.URL = url
	params.Payload = r
	if c.UA != "" {
		params.SetCustomHead("User-Agent", c.UA)
	}
	if modify != nil {
		params.AddFuncs(modify)
	}
	return params
}
func (c *Client) doHTTP(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	fn := func() error {
		resp, err = c.client.Do(req)
		return err
	}
	if err = c.retry(fn); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) reqHTTP(metod, url string, r io.Reader, modify request.ModifyFunc) (io.ReadCloser, error) {
	params := c.newRequsetParams(metod, url, r, modify)
	return params.FilterResponse(c.doHTTP)
}
