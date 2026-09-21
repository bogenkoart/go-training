package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/bogenkoart/go-training/internal/parallel"
)

func main() {
	fmt.Println("=== Задача 1: CallWithTimeout ===")

	fast := func() string { time.Sleep(50 * time.Millisecond); return "быстро" }
	slow := func() string { time.Sleep(300 * time.Millisecond); return "медленно" }

	res, err := parallel.CallWithTimeout(context.Background(), 100*time.Millisecond, fast)
	fmt.Printf("  быстрый вызов:   %q, err=%v\n", res, err)

	start := time.Now()
	res, err = parallel.CallWithTimeout(context.Background(), 100*time.Millisecond, slow)
	fmt.Printf("  медленный вызов: %q, err=%v, за %v\n",
		res, err, time.Since(start).Round(10*time.Millisecond))
	fmt.Printf("  это DeadlineExceeded: %v\n", errors.Is(err, context.DeadlineExceeded))

	parent, cancel := context.WithCancel(context.Background())
	go func() { time.Sleep(30 * time.Millisecond); cancel() }()
	start = time.Now()
	res, err = parallel.CallWithTimeout(parent, time.Second, slow)
	fmt.Printf("  отмена родителя: %q, err=%v, за %v\n",
		res, err, time.Since(start).Round(10*time.Millisecond))
	fmt.Printf("  это Canceled: %v\n", errors.Is(err, context.Canceled))

	time.Sleep(400 * time.Millisecond) // даём медленным вызовам досчитать до конца
	fmt.Printf("  горутин после всех вызовов: %d (ожидаем 1)\n", runtime.NumGoroutine())

	fmt.Println()
	fmt.Println("=== Задача 2: Heartbeat ===")

	ctx, stop := context.WithTimeout(context.Background(), 550*time.Millisecond)
	defer stop()

	start = time.Now()
	n := parallel.Heartbeat(ctx, 100*time.Millisecond, func() {})
	fmt.Printf("  ударов: %d (ожидаем 5), вернулся за %v\n",
		n, time.Since(start).Round(10*time.Millisecond))

	time.Sleep(200 * time.Millisecond)
	fmt.Printf("  горутин после остановки: %d (ожидаем 1)\n", runtime.NumGoroutine())
}
