package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	contentType     = "Content-Type"
	contentTypeJOSN = "application/json; charset=UTF-8"
)

// JSONRequest is a json req modifier
var JosnModify ModifyFunc = func(req *http.Request) {
	req.Header.Set(contentType, contentTypeJOSN)
}

// DecodeJSON decode resp body to container
// container should be a pointer type
func DecodeJSON(r io.ReadCloser, container interface{}) error {
	defer r.Close()
	if container == nil {
		return nil
	}
	return json.NewDecoder(r).Decode(container)
}

// MarshalJSON marshal payload to io.Reader
func MarshalJSON(payload interface{}) (io.Reader, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	return r, nil
}
