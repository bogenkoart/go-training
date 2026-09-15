package store

import (
	"errors"
	"fmt"
)

var _ Storage = (*MemoryStorage)(nil)
var _ Storage = (*LoggingStorage)(nil)
var ErrNotFound = errors.New("not found")

type Storage interface {
	Save(key string, value []byte) error
	Load(key string) ([]byte, error)
	Delete(key string) error
}

type MemoryStorage struct {
	data map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	memstor := make(map[string][]byte)
	return &MemoryStorage{
		data: memstor,
	}
}

func (m *MemoryStorage) Save(key string, value []byte) error {
	cop := make([]byte, len(value))
	copy(cop, value)
	m.data[key] = cop
	return nil
}

func (m *MemoryStorage) Load(key string) ([]byte, error) {
	if v, ok := m.data[key]; ok {
		cop := make([]byte, len(v))
		copy(cop, v)
		return cop, nil
	}
	return nil, ErrNotFound
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
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

// LoggingStorage встраивает интерфейс Storage.
//  Методы встроенного поля продвигаются к внешнему типу,
//  поэтому Load вызывается у того значения, которое лежит в поле Storage.
//  Save и Delete объявлены явно и перекрывают продвинутые версии.
