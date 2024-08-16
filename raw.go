package bst

// NewRaw creates a new Raw instance
func NewRaw[K, V any](compare func(K, K) int, b Backend[K, V], kvs ...KV[K, V]) *Raw[K, V] {
	s := makeRaw(compare, &sliceBackend[K, V]{}, kvs)
	return &s
}

func makeRaw[K, V any](compare func(K, K) int, b Backend[K, V], kvs []KV[K, V]) (s Raw[K, V]) {
	s.compare = compare
	s.b = b

	for _, kv := range kvs {
		s.Set(kv.Key, kv.Value)
	}

	return
}

// Raw is a Raw Stringset
type Raw[K, V any] struct {
	b       Backend[K, V]
	compare func(K, K) int
}

// Set will place a key
func (s *Raw[K, V]) Set(key K, value V) (err error) {
	var (
		index int
		match bool
	)

	if index, match, err = s.getIndex(key); err != nil {
		return
	}

	pair := makeKV(key, value)
	switch {
	case match:
		return s.b.Set(index, pair)
	case index == s.Len():
		return s.b.Append(pair)
	default:
		return s.b.InsertAt(index, pair)
	}
}

// Update will pass the existing value to the provided function and update the entry value with whatever is returned
func (s *Raw[K, V]) Update(key K, fn func(existing V) V) (err error) {
	var (
		index int
		match bool
	)

	if index, match, err = s.getIndex(key); err != nil {
		return
	}

	if !match {
		return makeErrorNotFound(key)
	}

	var kv KV[K, V]
	if kv, err = s.b.Get(index); err != nil {
		return
	}

	kv.Value = fn(kv.Value)
	return s.b.Set(index, kv)
}

// Get will retrieve a value for a given key
func (s *Raw[K, V]) Get(key K) (value V, err error) {
	var (
		index int
		match bool
	)

	if index, match, err = s.getIndex(key); err != nil {
		return
	}

	if !match {
		err = makeErrorNotFound(key)
		return
	}

	var kv KV[K, V]
	if kv, err = s.b.Get(index); err != nil {
		return
	}

	value = kv.Value
	return
}

// UsSet will remove a key
func (s *Raw[K, V]) RemoveAt(key K) (err error) {
	var (
		index int
		match bool
	)

	if index, match, err = s.getIndex(key); !match {
		return
	}

	return s.b.RemoveAt(index)
}

// Has will return if a key exists
func (s *Raw[K, V]) Has(key K) (has bool, err error) {
	_, has, err = s.getIndex(key)
	return
}

// Len will return the keys length
func (s *Raw[K, V]) Len() (n int) {
	return s.b.Len()
}

// Slice will return a slice with the copied contents
func (s *Raw[K, V]) Slice() (kvs []KV[K, V]) {
	return s.b.Slice()
}

// ForEach will iterate over all values
func (s *Raw[K, V]) ForEach(fn func(K, V) (end bool)) (ended bool) {
	return s.b.ForEach(func(kv KV[K, V]) (end bool) {
		return fn(kv.Key, kv.Value)
	})
}

// ForEach will iterate over all values
func (s *Raw[K, V]) Cursor(fn func(c *Cursor[K, V]) error) (err error) {
	var c Cursor[K, V]
	c.c = s.b.Cursor()
	c.r = s
	err = fn(&c)
	c.Close()
	return
}

func (s *Raw[K, V]) getKey(index int) (key K, err error) {
	var kv KV[K, V]
	if kv, err = s.b.Get(index); err != nil {
		return
	}

	return kv.Key, nil
}

func (s *Raw[K, V]) getIndex(key K) (index int, match bool, err error) {
	sz := s.Len()
	if sz == 0 {
		return
	}

	start := 0
	end := sz - 1
	index = sz / 2

	var endKey K
	if endKey, err = s.getKey(end); err != nil {
		return
	}

	if s.compare(endKey, key) == -1 {
		index = end + 1
		return
	}

	for {
		var ref K
		if ref, err = s.getKey(index); err != nil {
			return
		}

		compared := s.compare(key, ref)
		switch {
		case compared == 0:
			match = true
			return

		case start == end && compared == -1:
			// Use current index
			return
		case start == end && compared == 1:
			index++
			return

		case compared == -1:
			if index == start {
				return
			}

			end = index - 1
		case compared == 1:
			if index == end {
				return
			}

			start = index + 1
		}

		index = (start + end) / 2
	}
}
