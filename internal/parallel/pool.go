package parallel

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// ProcessLimited обрабатывает items, выполняя не более workers задач одновременно.
// Порядок результатов соответствует порядку items.
func ProcessLimited(items []string, workers int) []string {
	if workers <= 0 {
		log.Printf("неверный параметр: workers <= 0\n")
		return nil
	}

	results := make([]string, len(items))
	sem := make(chan struct{}, workers)
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		go func(item string, i int) {
			sem <- struct{}{}
			defer func() { <-sem }()
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					log.Println("panic: process")
				}
			}()
			results[i] = process(item)
		}(item, i)
	}

	wg.Wait()
	return results
}

func process(item string) string {
	time.Sleep(50 * time.Millisecond)
	return fmt.Sprintf("done: %s", item)
}
