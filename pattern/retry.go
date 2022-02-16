package pattern

import (
	"context"
	"time"
)

// Task is a function that can be retried.
type Task func() error

func Retry(ctx context.Context, task Task, retryCount int) error {
	for i := 0; i < retryCount; i++ {
		if err := task(); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(2 << (i + 1) * time.Second)
		}
	}
	return task()
}
