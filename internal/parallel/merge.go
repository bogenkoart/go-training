package parallel

import "sync"

// Merge объединяет несколько каналов в один.
// Результирующий канал закрывается, когда закрыты все входные.
func Merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
