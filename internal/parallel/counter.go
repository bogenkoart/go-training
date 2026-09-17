package parallel

import (
	"sync"
	"sync/atomic"
)

const (
	goroutines = 100
	increments = 100
)

func CounterRacy() int {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	return counter
}
func CounterMutex() int {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return counter
}

func CounterAtomic() int64 {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter.Add(1)
			}
		}()
	}

	wg.Wait()
	return counter.Load()
}
