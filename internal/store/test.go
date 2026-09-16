package store

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