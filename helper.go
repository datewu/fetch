package fetch

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	contentType     = "Content-Type"
	contentTypeJOSN = "application/json; charset=UTF-8"
)

type reqModify func(*http.Request)

var josnModify reqModify = func(req *http.Request) {
	req.Header.Set(contentType, contentTypeJOSN)
}

func (c *Client) init() {
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

func (c *Client) reqHTTP(metod, url string, r io.Reader, modify reqModify) (io.ReadCloser, error) {
	req, err := http.NewRequest(metod, url, r)
	if err != nil {
		return nil, err
	}
	if modify != nil {
		modify(req)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	return resp.Body, nil
}

// container should be a pointer type
func respJSON(r io.ReadCloser, container interface{}) error {
	defer r.Close()
	return json.NewDecoder(r).Decode(container)
}

func reqJSON(payload interface{}) (io.Reader, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	return r, nil
}
