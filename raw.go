package bst

import "fmt"

// NewRaw a new Raw instance
func NewRaw[K, V any](compare func(K, K) int, kvs ...KV[K, V]) *Raw[K, V] {
	s := makeRaw[K, V](compare, &sliceBackend[K, V]{})
	return &s
}

func makeRaw[K, V any](compare func(K, K) int, b Backend[K, V]) (s Raw[K, V]) {
	s.compare = compare
	s.b = b
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
	if match {
		s.b.Set(index, pair)
		return
	}

	s.b.InsertAt(index, pair)
	return
}

// Set will place a key
func (s *Raw[K, V]) Update(key K, fn func(V) V) (err error) {
	var (
		index int
		match bool
	)

	if index, match, err = s.getIndex(key); err != nil {
		return
	}

	if !match {
		err = fmt.Errorf("entry of <%v> was not found", key)
		return
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
		err = fmt.Errorf("entry of <%v> was not found", key)
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

// Len will return the keys length
func (s *Raw[K, V]) Slice() (kvs []KV[K, V]) {
	return s.b.Slice()
}

// Len will return the keys length
func (s *Raw[K, V]) ForEach(fn func(KV[K, V]) (end bool)) (ended bool) {
	return s.b.ForEach(fn)
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
			end = index - 1
		case compared == 1:
			start = index + 1
		}

		index = (start + end) / 2
	}
}
