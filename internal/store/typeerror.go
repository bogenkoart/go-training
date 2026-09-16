package store

import "fmt"

// KeyError описывает ошибку операции над конкретным ключом.
type KeyError struct {
	Op  string // какая операция: "load", "delete", "save"
	Key string // над каким ключом
	Err error  // исходная причина
}

func (e *KeyError) Error() string {
	return fmt.Sprintf("%s %q: %v", e.Op, e.Key, e.Err)
}

func (e *KeyError) Unwrap() error {
	return e.Err
}
