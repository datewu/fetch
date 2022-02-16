package fetch

import (
	"context"

	"github.com/datewu/fetch/client"
)

var defaultClient = client.DefaultClient(context.Background())

// QuickGet is a shortcut for defaultClient.Get(url, container)
// container should be a pointer type
func QuickGet(url string, container interface{}) error {
	return defaultClient.Get(url, container)
}

// QuickPost is a shortcut for defaultClient.Post(url, payload, container)
// container should be a pointer type
func QuickPost(url string, data interface{}, container interface{}) error {
	return defaultClient.Post(url, data, container)
}
