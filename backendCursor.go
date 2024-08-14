package bst

type BackendCursor[T any] interface {
	Seek(int) (T, bool)
	Next() (T, bool)
	Prev() (T, bool)
	Close()
}
