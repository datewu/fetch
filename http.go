package fetch

// container should be a pointer type
func (c *Client) QueryJSON(method, url string, container interface{}) error {
	body, err := c.reqHTTP(method, url, nil, nil)
	if err != nil {
		return err
	}
	defer body.Close()
	return respJSON(body, container)
}

// container should be a pointer type
func (c *Client) CreatJSON(method, url string, payload, container interface{}) error {
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
