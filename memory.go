package bst

func NewMemory[K any, V any](initialCapacity int64, compare func(K, K) int, kvs ...KV[K, V]) (out *Memory[K, V], err error) {
	var m Memory[K, V]
	m.s = make(sliceBackend[K, V], 0, initialCapacity)
	m.Raw = makeRaw(compare, &m.s, kvs)
	out = &m
	return
}

type Memory[K any, V any] struct {
	s sliceBackend[K, V]
	Raw[K, V]
}
