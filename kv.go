package bst

func makeKV(key string, value interface{}) (kv KV) {
	kv.Key = key
	kv.Value = value
	return
}

// KV represents a Key/Value pair
type KV struct {
	Key   string
	Value interface{}
}
