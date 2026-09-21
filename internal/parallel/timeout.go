package parallel

import (
	"context"
	"time"
)

// CallWithTimeout вызывает slow и ждёт результат не дольше timeout.
// Если ctx отменят раньше — возвращается немедленно.
func CallWithTimeout(ctx context.Context, timeout time.Duration, slow func() string) (string, error) {
	callctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		ch <- slow()
	}()

	select {
	case v := <-ch:
		return v, nil
	case <-callctx.Done():
		return "", callctx.Err()
	}
}

// Heartbeat вызывает beat каждые interval, пока не отменят ctx.
// Возвращает количество выполненных вызовов.
func Heartbeat(ctx context.Context, interval time.Duration, beat func()) int {
	if interval <= 0 {
		return 0
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ctx.Done():
			return count
		case <-ticker.C:
			if ctx.Err() != nil {
				return count
			}
			beat()
			count++
		}
	}
}
