package bst

// NewStore a new Store instance
func NewStore[V any](kvs ...KV[string, V]) *Store[V] {
	s := makeStore[V](kvs)
	return &s
}

// NewStore a new Store instance
func makeStore[V any](kvs []KV[string, V]) (s Store[V]) {
	s.Raw = makeRaw[string, V](compareString, kvs)
	return
}

type Store[V any] struct {
	Raw[string, V]
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
