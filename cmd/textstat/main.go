package main

import (
	"fmt"
	"time"

	"github.com/bogenkoart/go-training/internal/parallel"
)

func main() {
	fmt.Println("=== Задача 1: CountWordsParallel ===")
	texts := []string{
		"привет мир мир",
		"мир go go",
		"go привет",
	}
	// пять прогонов подряд: результат обязан совпасть пять раз из пяти
	for i := 1; i <= 5; i++ {
		fmt.Printf("  прогон %d: %v\n", i, parallel.CountWordsParallel(texts))
	}
	fmt.Printf("  пустой вход: %v\n", parallel.CountWordsParallel(nil))

	fmt.Println()
	fmt.Println("=== Задача 2: счётчики, по 5 прогонов ===")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  racy=%-6d mutex=%-6d atomic=%-6d\n",
			parallel.CounterRacy(),
			parallel.CounterMutex(),
			parallel.CounterAtomic(),
		)
	}

	fmt.Println()
	fmt.Println("=== Время, среднее из 10 прогонов ===")
	fmt.Printf("  racy:   %v\n", measure(func() { parallel.CounterRacy() }))
	fmt.Printf("  mutex:  %v\n", measure(func() { parallel.CounterMutex() }))
	fmt.Printf("  atomic: %v\n", measure(func() { parallel.CounterAtomic() }))
}

// measure прогоняет f десять раз и возвращает среднее время одного прогона.
func measure(f func()) time.Duration {
	const runs = 10
	start := time.Now()
	for i := 0; i < runs; i++ {
		f()
	}
	return time.Since(start) / runs
}
