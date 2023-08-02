package bst

// NewKeys a new Keys instance
func NewKeys(keys ...string) *Keys {
	var k Keys
	k.Store = makeStore(nil)
	for _, key := range keys {
		k.Store.Set(key, struct{}{})
	}

	return &k
}

// Keys is a Keys Stringset
type Keys struct {
	Store
}

// Set will place a key
func (k *Keys) Set(key string) {
	k.Store.Set(key, struct{}{})
}

// Len will return the keys length
func (k *Keys) ForEach(fn func(string) (end bool)) (ended bool) {
	return k.Store.ForEach(func(key string, _ interface{}) (end bool) {
		return fn(key)
	})
}
