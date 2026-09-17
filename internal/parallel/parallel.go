package parallel

import (
	"sync"

	"github.com/bogenkoart/go-training/internal/textstat"
)

func CountWordsParallel(texts []string) map[string]int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	resultCnt := make(map[string]int)

	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			words := textstat.Words(text)
			wordCnt := textstat.Frequency(words)
			mu.Lock()
			for key, val := range wordCnt {
				resultCnt[key] += val
			}
			mu.Unlock()
		}(text)
	}
	wg.Wait()
	return resultCnt
}
