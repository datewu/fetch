package fetch

import "net/http"

// Get is a shortcut for c.getJSON(http.MethodGet, url, container)
func (c *Client) Get(url string, container interface{}) error {
	return c.getJSON(http.MethodGet, url, container)
}

// Delete is a shortcut for c.getJSON(http.MethodDelete, url, container)
func (c *Client) Delete(url string, container interface{}) error {
	return c.getJSON(http.MethodDelete, url, container)
}

// Post is a shortcut for c.setJSON(http.MethodPost, url, payload, container)
func (c *Client) Post(method, url string, payload, container interface{}) error {
	return c.setJSON(http.MethodPost, url, payload, container)
}

// Put is a shortcut for c.setJSON(http.MethodPut, url, payload, container)
func (c *Client) Put(method, url string, payload, container interface{}) error {
	return c.setJSON(http.MethodPut, url, payload, container)
}

// Patch is a shortcut for c.setJSON(http.MethodPatch, url, payload, container)
func (c *Client) Patch(method, url string, payload, container interface{}) error {
	return c.setJSON(http.MethodPatch, url, payload, container)
}

// container should be a pointer type
func (c *Client) getJSON(method, url string, container interface{}) error {
	body, err := c.reqHTTP(method, url, nil, nil)
	if err != nil {
		return err
	}
	defer body.Close()
	return respJSON(body, container)
}

// container should be a pointer type
func (c *Client) setJSON(method, url string, payload, container interface{}) error {
	r, err := reqJSON(payload)
	if err != nil {
		return err
	}
	body, err := c.reqHTTP(method, url, r, josnModify)
	if err != nil {
		return err
	}
	return respJSON(body, container)
}
