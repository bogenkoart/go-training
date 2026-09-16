package store

import (
	"errors"
	"fmt"
	"runtime/debug"
)

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

type PanicStorage struct{ Storage } // встраиваем интерфейс — методы продвинутся

func (p PanicStorage) Load(key string) ([]byte, error) {
	panic("хранилище сломалось")
}

func (p PanicStorage) Save(key string, value []byte) error {
	panic("хранилище сломалось")
}

func (p PanicStorage) Delete(key string) error {
	panic("хранилище сломалось")
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

func (m *MemoryStorage) Load(key string) ([]byte, *KeyError) {
	if v, ok := m.data[key]; ok {
		cop := make([]byte, len(v))
		copy(cop, v)
		return cop, nil
	}
	return nil, &KeyError{
		Op:  "load",
		Key: key,
		Err: ErrNotFound,
	}
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

// LoggingStorage встраивает интерфейс Storage.
//  Методы встроенного поля продвигаются к внешнему типу,
//  поэтому Load вызывается у того значения, которое лежит в поле Storage.
//  Save и Delete объявлены явно и перекрывают продвинутые версии.

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

// SafeLoad вызывает s.Load и превращает любую панику внутри реализации
// в обычную ошибку, не роняя программу.
func SafeLoad(s Storage, key string) (data []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("safe load %q паника: %v, %v", key, r, debug.Stack())
		}
	}()
	data, err = s.Load(key)
	return data, err
}

// потому что  если мы не укажем именнованным err, то мы не сможем передать ошибку далше и функция вернёт корректное срабатывание в случае паники. Мы просто не заметим панику
