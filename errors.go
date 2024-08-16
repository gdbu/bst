package bst

import "fmt"

func makeErrorNotFound[T any](key T) (e ErrorNotFound[T]) {
	e.key = key
	return
}

type ErrorNotFound[T any] struct {
	key T
}

func (e ErrorNotFound[T]) Key() T {
	return e.key
}

func (e ErrorNotFound[T]) Error() string {
	return fmt.Sprintf("entry of <%v> was not found", e.key)
}

func IsEntryNotFound[T any](err error) (ok bool) {
	if err == nil {
		return
	}

	_, ok = err.(ErrorNotFound[T])
	return
}
