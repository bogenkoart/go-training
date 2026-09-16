package store

package store

import "testing"

// тип неэкспортируемый — тест лежит в том же пакете и видит его
type panicStorage struct{ Storage }

func (p panicStorage) Load(key string) ([]byte, error) {
	panic("хранилище сломалось")
}

func TestSafeLoad_Panic(t *testing.T) {
	data, err := SafeLoad(panicStorage{}, "k")
	if err == nil {
		t.Fatal("ожидали ошибку, получили nil")
	}
	if data != nil {
		t.Errorf("ожидали nil, получили %v", data)
	}
}