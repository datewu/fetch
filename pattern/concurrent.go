package pattern

import (
	"errors"
	"sync"
)

// Fetcher use with closure
type Fetcher interface {
	Fetch() (interface{}, error)
}

// ConcurrentFirstOrErr concurrent fetch the fetcher list
// return the first successful result or err when all fetchers failed.
func ConcurrentFirstOrErr(fs []Fetcher) (interface{}, error) {
	if len(fs) == 0 {
		return nil, errors.New("empty fetcher list")
	}
	var wg sync.WaitGroup
	wg.Add(len(fs))
	done := make(chan struct{})
	resStream := make(chan interface{})
	errStream := make(chan error)
	for _, f := range fs {
		go func(lookup Fetcher) {
			defer wg.Done()
			res, err := lookup.Fetch()
			if err != nil {
				select {
				case <-done:
				case errStream <- err:
				}
				return
			}
			select {
			case <-done:
			case resStream <- res:
			}
		}(f)
	}
	go func() {
		wg.Wait()
		close(resStream)
		close(errStream)
	}()
	out := <-resStream
	close(done)
	if out != nil {
		return out, nil
	}
	select {
	case err := <-errStream:
		return nil, err
	default:
		return nil, errors.New("no reust")
	}
}
