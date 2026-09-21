package safe

import (
	"sync"
	"sync/atomic"
)

type RateLimiter struct {
	mu       sync.RWMutex
	limiter  map[string]int
	totalcnt atomic.Int64
	limit    int
}

func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{limiter: make(map[string]int), limit: limit}
}

// Allow возвращает true, если пользователь ещё не исчерпал limit запросов,
// и засчитывает запрос. Если лимит исчерпан — false, запрос не засчитывается.
func (l *RateLimiter) Allow(user string) bool {
	l.totalcnt.Add(1)
	l.mu.Lock()
	defer l.mu.Unlock()
	if v := l.limiter[user]; v < l.limit {
		l.limiter[user]++
		return true
	}
	return false
}

// Count возвращает, сколько запросов пользователя уже засчитано.
func (l *RateLimiter) Count(user string) int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.limiter[user]
}

// Total возвращает, сколько раз вообще вызывался Allow — разрешённых и нет.
func (l *RateLimiter) Total() int64 {
	return l.totalcnt.Load()
}
