package fetch

import (
	"context"

	"github.com/datewu/fetch/client"
)

// QuickClient is a quick client which persistently use the default client
var QuickClient = client.DefaultClient(context.Background())

// QuickGet quick grab something
// container should be a pointer type
func QuickGet(url string, container interface{}) error {
	cli := client.DefaultClient(context.Background())
	return cli.Get(url, container)
}

// QuickPost quick post something
// container should be a pointer type
func QuickPost(url string, data interface{}, container interface{}) error {
	cli := client.DefaultClient(context.Background())
	return cli.Post(url, data, container)
}
