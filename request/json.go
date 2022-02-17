package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
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
	t := reflect.TypeOf(container)
	if t.Kind() != reflect.Ptr {
		return errors.New("container should be a pointer type")
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
