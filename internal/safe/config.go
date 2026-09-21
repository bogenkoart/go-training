package safe

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	configOnce sync.Once
	config     *Config
)

type Config struct {
	Addr    string
	Workers int
}

var loads atomic.Int64

// loadConfig имитирует медленное чтение файла и считает свои вызовы.
func loadConfig() *Config {
	loads.Add(1)
	time.Sleep(50 * time.Millisecond)
	return &Config{Addr: ":8080", Workers: 4}
}

// LoadCount — сколько раз вызывалась loadConfig. Для проверки.
func LoadCount() int64 { return loads.Load() }

// GetConfig возвращает конфиг, загружая его при первом обращении.
func GetConfig() *Config {
	configOnce.Do(func() {
		config = loadConfig()
	})
	return config
}
