package bst

// New a new Store instance
func New[V any](kvs ...KV[string, V]) *Store[V] {
	s := makeStore(&sliceBackend[string, V]{}, kvs)
	return &s
}

// NewStore a new Store instance
func makeStore[V any](b Backend[string, V], kvs []KV[string, V]) (s Store[V]) {
	s.Raw = makeRaw(compareString, b, kvs)
	return
}

type Store[V any] struct {
	Raw[string, V]
}

// Len will return the keys length
func (s *Store[V]) ForEach(fn func(string, V) (end bool)) (ended bool) {
	return s.Raw.ForEach(fn)
}

func compareString(a, b string) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	default:
		return 1
	}
}
