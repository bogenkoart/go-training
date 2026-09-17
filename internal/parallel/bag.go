package parallel

import (
	"sync"
	"sync/atomic"
)

var Iter int = 100
var Iterinc int = 100

func CounterRacy() int {
	var counter int

	for i := 0; i < Iter; i++ {
		go func() {
			for j := 0; j < Iterinc; j++ {
				counter++
			}
		}()
	}
	return counter
}
func CounterMutex() int {
	var counter int
	var mu sync.Mutex

	for i := 0; i < Iter; i++ {
		go func() {
			for j := 0; j < Iterinc; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	return counter
}

func CounterAtomic() int64 {
	var counter atomic.Int64

	for i := 0; i < Iter; i++ {
		go func() {
			for j := 0; j < Iterinc; j++ {
				counter.Add(1)
			}
		}()
	}
	return counter.Load()
}
