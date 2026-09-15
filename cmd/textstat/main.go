package main

import (
	"fmt"

	"github.com/bogenkoart/go-training/internal/store"
)

func main() {
	var _ store.Storage = (*store.MemoryStorage)(nil) // вот здесь потом поясни пожалуйста
	var _ store.Storage = (*store.LoggingStorage)(nil)
	s := store.NewMemoryStorage()
	data := []byte("hello")
	s.Save("k", data)
	data[0] = 'X'
	got, _ := s.Load("k")
	fmt.Println(string(got))
	s1, _ := s.Load("k")
	s1[0] = 'X'
	got1, _ := s.Load("k")
	fmt.Println(string(got1))
	fmt.Println(s)
	s.Delete("k")
	got1, err := s.Load("k")
	fmt.Println(err)
	fmt.Println(s)
	fmt.Println()

	l := store.LoggingStorage{Storage: store.NewMemoryStorage()}
	data1 := []byte("hello")
	l.Save("k", data1)
	m, _ := l.Load("k")
	fmt.Println(string(m))
	l.Delete("k")
	l.Delete("k")
}
