package fetch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuick(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, fetch")
	}))
	defer ts.Close()
	t.Log("test server url", ts.URL)
	if err := QuickGet(ts.URL, nil); err != nil {
		t.Error(err)
	}
	todo := make(map[string]interface{})
	err := QuickGet("https://jsonplaceholder.typicode.com/todos/1", &todo)
	if err != nil {
		t.Error(err)
	}
	if len(todo) == 0 {
		t.Error("todo is empty")
	}
}
func TestPost(t *testing.T) {
	todo := map[string]interface{}{
		"userId":    1,
		"id":        1,
		"title":     "delectus aut autem",
		"completed": false,
	}

	resp := make(map[string]interface{})
	err := QuickPost("https://jsonplaceholder.typicode.com/todos", todo, &resp)
	if err != nil {
		t.Error(err)
	}
	if len(resp) == 0 {
		t.Error("resp is empty")
	}
}
