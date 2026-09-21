package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/bogenkoart/go-training/internal/safe"
)

func main() {
	fmt.Println("=== Задача 1: RateLimiter ===")
	l := safe.NewRateLimiter(50)

	var allowed atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				if l.Allow("alice") {
					allowed.Add(1)
				}
			}
		}()
	}
	wg.Wait()

	fmt.Printf("  разрешено alice: %d (ожидаем ровно 50)\n", allowed.Load())
	fmt.Printf("  Count(alice):    %d (ожидаем 50)\n", l.Count("alice"))
	fmt.Printf("  Allow(bob):      %v (ожидаем true — лимит у каждого свой)\n", l.Allow("bob"))
	fmt.Printf("  Total():         %d (ожидаем 1001 — все вызовы Allow)\n", l.Total())

	fmt.Println()
	fmt.Println("=== Задача 2: GetConfig ===")
	ptrs := make([]*safe.Config, 100)
	var wg2 sync.WaitGroup
	for i := range ptrs {
		wg2.Add(1)
		go func(i int) {
			defer wg2.Done()
			ptrs[i] = safe.GetConfig()
		}(i)
	}
	wg2.Wait()

	same := true
	for _, p := range ptrs {
		if p != ptrs[0] {
			same = false
		}
	}
	fmt.Printf("  загрузок: %d (ожидаем 1)\n", safe.LoadCount())
	fmt.Printf("  у всех один и тот же объект: %v\n", same)
	fmt.Printf("  конфиг: %+v\n", *ptrs[0])
}
