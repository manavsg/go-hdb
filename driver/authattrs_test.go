package driver

import (
	"sync"
	"testing"
)

// test if concurrent refresh would deadlock.
func testConcurrentRefresh(t *testing.T) {
	const numConcurrent = 100

	attrs := &authAttrs{}
	attrs.SetRefreshPassword(func() (string, bool) { return "", true })
	attrs.SetRefreshToken(func() (string, bool) { return "", true })
	attrs.SetRefreshClientCert(func() ([]byte, []byte, bool) { return nil, nil, true })

	wg := new(sync.WaitGroup)
	wg.Add(numConcurrent)
	start := make(chan struct{})
	for i := 0; i < numConcurrent; i++ {
		go func(start <-chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			<-start
			attrs.refresh() //nolint:errcheck
		}(start, wg)
	}
	// start refresh concurrently
	close(start)
	// wait for all go routines to end
	wg.Wait()
}

func TestAuthAttrs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fct  func(t *testing.T)
	}{
		{"testConcurrentRefresh", testConcurrentRefresh},
	}

	for _, test := range tests {
		test := test // new test to run in parallel

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.fct(t)
		})
	}
}
