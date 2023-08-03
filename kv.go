package bst

func makeKV[K, V any](key K, value V) (kv KV[K, V]) {
	kv.Key = key
	kv.Value = value
	return
}

// KV represents a Key/Value pair
type KV[K, V any] struct {
	Key   K
	Value V
}
