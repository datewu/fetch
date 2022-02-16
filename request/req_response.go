package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type ModifyFunc func(*http.Request)

type ReqParams struct {
	ctx           context.Context
	Method        string
	URL           string
	Payload       io.Reader
	customHeaders map[string]string
	modifyFuncs   []ModifyFunc
}

func NewReqParams(ctx context.Context) *ReqParams {
	params := &ReqParams{
		ctx:           ctx,
		customHeaders: make(map[string]string),
		modifyFuncs:   make([]ModifyFunc, 0),
	}
	return params
}

func (p *ReqParams) SetCustomHead(k, v string) {
	p.customHeaders[k] = v
}
func (p *ReqParams) AddFuncs(fn ModifyFunc) {
	p.modifyFuncs = append(p.modifyFuncs, fn)
}
func (p *ReqParams) GetRequest() (*http.Request, error) {
	req, err := http.NewRequestWithContext(p.ctx, p.Method, p.URL, p.Payload)
	if err != nil {
		return nil, err
	}
	for k, v := range p.customHeaders {
		req.Header.Set(k, v)
	}
	for _, v := range p.modifyFuncs {
		v(req)
	}
	return req, nil
}

type ResponseFunc func(req *http.Request) (*http.Response, error)

func (p *ReqParams) FilterResponse(do ResponseFunc) (io.ReadCloser, error) {
	req, err := p.GetRequest()
	if err != nil {
		return nil, err
	}
	resp, err := do(req)
	if err != nil {
		return nil, err
	}
	// fn := func() error {
	// 	if c.client == nil {
	// 		c.setDefaultCli()
	// 	}
	// 	resp, err = c.client.Do(req)
	// 	return err
	// }
	// if err = c.retry(fn); err != nil {
	// 	return nil, err
	// }
	if resp.StatusCode >= http.StatusBadRequest {
		resp.Body.Close()
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	return resp.Body, nil

}
