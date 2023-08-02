package bst

func makeKV[T any](key string, value T) (kv KV[T]) {
	kv.Key = key
	kv.Value = value
	return
}

// KV represents a Key/Value pair
type KV[T any] struct {
	Key   string
	Value T
}
