package fetch

import (
	"testing"
)

func TestGet(t *testing.T) {
	if err := QuickGet("https://jd.com", nil); err != nil {
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
