package main

import (
	"errors"
	"fmt"

	"github.com/bogenkoart/go-training/internal/store"
)

func main() {
	s := store.NewMemoryStorage()
	_, err := s.Load("user:17")

	// 1. Сентинельная ошибка находится СКВОЗЬ твою обёртку
	fmt.Println(errors.Is(err, store.ErrNotFound)) // ожидаем: true

	// 2. Из ошибки достаются детали
	var ke *store.KeyError
	if errors.As(err, &ke) {
		fmt.Println(ke.Op, ke.Key) // ожидаем: load user:17
	}

	// 3. Текст содержит и операцию, и ключ
	fmt.Println(err) // ожидаем: load "user:17": not found
}
