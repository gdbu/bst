package bst

// NewKeys a new Keys instance
func NewKeys(keys ...string) *Keys {
	var k Keys
	k.s = makeStore[struct{}](nil)
	for _, key := range keys {
		k.s.Set(key, struct{}{})
	}

	return &k
}

// Keys is a Keys Stringset
type Keys struct {
	s Store[struct{}]
}

// Set will place a key
func (k *Keys) Set(key string) {
	k.s.Set(key, struct{}{})
}

// Unset removes a key
func (k *Keys) Unset(key string) {
	k.s.Unset(key)
}

// Has determines if a key exists
func (k *Keys) Has(key string) bool {
	return k.s.Has(key)
}

// Len returns the length of the underlying keys
func (k *Keys) Len() (n int) {
	return k.s.Len()
}

// Len will return the keys length
func (k *Keys) ForEach(fn func(string) (end bool)) (ended bool) {
	return k.s.ForEach(func(key string, _ struct{}) (end bool) {
		return fn(key)
	})
}

// Has determines if a key exists
func (k *Keys) Slice() (out []string) {
	out = make([]string, 0, k.s.Len())
	k.ForEach(func(str string) (end bool) {
		out = append(out, str)
		return
	})

	return
}
