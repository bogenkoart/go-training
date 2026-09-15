package store

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")
var ErrSave = errors.New("такой ключ уже существует")

type Storage interface {
	Save(key string, value []byte) error
	Load(key string) ([]byte, error)
	Delete(key string) error
}

type MemoryStorage struct {
	memstor map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	memstor := make(map[string][]byte)
	return &MemoryStorage{
		memstor: memstor,
	}
}

func (m *MemoryStorage) Save(key string, value []byte) error {
	cop := make([]byte, len(value))
	copy(cop, value)
	if _, ok := m.memstor[key]; !ok {
		m.memstor[key] = cop
		return nil
	}
	return ErrSave
}

func (m *MemoryStorage) Load(key string) ([]byte, error) {
	if v, ok := m.memstor[key]; ok {
		cop := make([]byte, len(v))
		copy(cop, v)
		return cop, nil
	}
	return []byte{}, ErrNotFound
}

func (m *MemoryStorage) Delete(key string) error {
	if _, ok := m.memstor[key]; ok {
		delete(m.memstor, key)
		return nil
	}
	return ErrNotFound
}

type LoggingStorage struct {
	Storage // встроенный ИНТЕРФЕЙС, не структура
}

func (l *LoggingStorage) Save(key string, value []byte) error {
	fmt.Println("Добавление элемента")
	err := l.Storage.Save(key, value)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (l *LoggingStorage) Delete(key string) error {
	fmt.Println("Удаление элемента")
	err := l.Storage.Delete(key)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
