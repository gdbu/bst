package bst

type Backend[K, V any] interface {
	Get(index int) KV[K, V]
	Set(int, V)
	Unset(int)
	InsertAt(int, KV[K, V])
	Len() int
	Slice() []KV[K, V]
	ForEach(fn func(K, V) (end bool)) (ended bool)
}
