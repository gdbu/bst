package bst

type Backend[K, V any] interface {
	Get(index int) (KV[K, V], error)
	Set(int, KV[K, V]) error
	InsertAt(int, KV[K, V]) error
	RemoveAt(int) error
	Len() int
	Slice() []KV[K, V]
	ForEach(fn func(KV[K, V]) (end bool)) (ended bool)
}
