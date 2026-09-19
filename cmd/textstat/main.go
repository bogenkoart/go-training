package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/bogenkoart/go-training/internal/parallel"
)

func main() {
	fmt.Println("=== Задача 1: ProcessLimited ===")

	items := make([]string, 12)
	for i := range items {
		items[i] = fmt.Sprintf("item-%02d", i)
	}

	start := time.Now()
	res := parallel.ProcessLimited(items, 3)
	elapsed := time.Since(start)

	fmt.Printf("  результатов: %d\n", len(res))
	fmt.Printf("  первые три:  %v\n", res[:3])
	fmt.Printf("  порядок ок:  %v\n", checkOrder(items, res))
	fmt.Printf("  время:       %v\n", elapsed)
	fmt.Println("  ожидаем ~4 волны по 50мс = ~200мс; если ~50мс — лимит не работает")

	fmt.Println()
	fmt.Println("=== Задача 2: Merge ===")

	before := runtime.NumGoroutine()

	a := gen(1, 2, 3)
	b := gen(10, 20)
	c := gen(100)

	var sum int64
	count := 0
	for v := range parallel.Merge(a, b, c) {
		atomic.AddInt64(&sum, int64(v))
		count++
	}
	fmt.Printf("  получено значений: %d (ожидаем 6)\n", count)
	fmt.Printf("  сумма: %d (ожидаем 136)\n", sum)

	// пустой вызов
	empty := 0
	for range parallel.Merge() {
		empty++
	}
	fmt.Printf("  Merge() без аргументов вернул значений: %d (ожидаем 0)\n", empty)

	time.Sleep(100 * time.Millisecond) // даём горутинам доиграть
	after := runtime.NumGoroutine()
	fmt.Printf("  горутин было %d, стало %d — утечки %v\n",
		before, after, map[bool]string{true: "нет", false: "ЕСТЬ"}[after <= before])
}

// gen возвращает канал, отдающий переданные значения и закрывающийся после.
func gen(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range vals {
			ch <- v
		}
	}()
	return ch
}

// checkOrder проверяет, что results[i] получен из items[i].
func checkOrder(items, results []string) bool {
	if len(items) != len(results) {
		return false
	}
	for i := range items {
		if results[i] != "done:"+items[i] {
			return false
		}
	}
	return true
}
