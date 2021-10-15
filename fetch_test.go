package fetch

import (
	"context"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	p := context.Background()
	ctx, cancel := context.WithTimeout(p, 2*time.Second)
	defer cancel()
	cli := DefaultClient(ctx)
	if err := cli.Get("https://jd.com", nil); err != nil {
		t.Error(err)
	}
	todo := make(map[string]interface{})
	err := cli.Get("https://jsonplaceholder.typicode.com/todos/1", &todo)
	if err != nil {
		t.Error(err)
	}
	if len(todo) == 0 {
		t.Error("todo is empty")
	}
}
func TestPost(t *testing.T) {
	p := context.Background()
	ctx, cancel := context.WithTimeout(p, 2*time.Second)
	defer cancel()
	cli := DefaultClient(ctx)
	todo := map[string]interface{}{
		"userId":    1,
		"id":        1,
		"title":     "delectus aut autem",
		"completed": false,
	}

	resp := make(map[string]interface{})
	err := cli.Post("https://jsonplaceholder.typicode.com/todos", todo, &resp)
	if err != nil {
		t.Error(err)
	}
	if len(resp) == 0 {
		t.Error("resp is empty")
	}
}
