package bst

import (
	"sort"
)

// NewRaw a new Raw instance
func NewRaw[K, V any](compare func(K, K) int, kvs ...KV[K, V]) *Raw[K, V] {
	s := makeRaw[K, V](compare, kvs)
	return &s
}

func makeRaw[K, V any](compare func(K, K) int, kvs []KV[K, V]) (s Raw[K, V]) {
	sz := len(kvs)
	if sz == 0 {
		sz = 8
	}

	s.compare = compare
	s.kvs = make([]KV[K, V], len(kvs), sz)
	copy(s.kvs, kvs)
	sort.Slice(s.kvs, func(i, j int) (less bool) {
		return s.compare(s.kvs[i].Key, s.kvs[j].Key) == -1
	})

	return
}

// Raw is a Raw Stringset
type Raw[K, V any] struct {
	kvs     []KV[K, V]
	compare func(K, K) int
}

// Set will place a key
func (s *Raw[K, V]) Set(key K, value V) {
	index, match := s.getIndex(key)
	if match {
		s.kvs[index].Value = value
		return
	}

	pair := makeKV(key, value)
	first := s.kvs[:index]
	second := append([]KV[K, V]{pair}, s.kvs[index:]...)
	s.kvs = append(first, second...)
}

// Set will place a key
func (s *Raw[K, V]) Update(key K, fn func(V) V) (success bool) {
	index, match := s.getIndex(key)
	if !match {
		return false
	}

	s.kvs[index].Value = fn(s.kvs[index].Value)
	return true
}

// Get will retrieve a value for a given key
func (s *Raw[K, V]) Get(key K) (value V, has bool) {
	var index int
	if index, has = s.getIndex(key); !has {
		return
	}

	value = s.kvs[index].Value
	return
}

// UsSet will remove a key
func (s *Raw[K, V]) Unset(key K) {
	index, match := s.getIndex(key)
	if !match {
		return
	}

	first := s.kvs[:index]
	second := s.kvs[index+1:]
	s.kvs = append(first, second...)
}

// Has will return if a key exists
func (s *Raw[K, V]) Has(key K) (has bool) {
	_, has = s.getIndex(key)
	return
}

// Len will return the keys length
func (s *Raw[K, V]) Len() (n int) {
	return len(s.kvs)
}

// Len will return the keys length
func (s *Raw[K, V]) Slice() (kvs []KV[K, V]) {
	kvs = make([]KV[K, V], len(s.kvs))
	copy(kvs, s.kvs)
	return
}

// Len will return the keys length
func (s *Raw[K, V]) ForEach(fn func(K, V) (end bool)) (ended bool) {
	for _, kv := range s.kvs {
		if ended = fn(kv.Key, kv.Value); ended {
			return
		}
	}

	return
}

func (s *Raw[K, V]) getKey(index int) K {
	return s.kvs[index].Key
}

func (s *Raw[K, V]) getIndex(key K) (index int, match bool) {
	sz := s.Len()
	if sz == 0 {
		return
	}

	start := 0
	end := sz - 1
	index = sz / 2

	if s.compare(s.getKey(end), key) == -1 {
		index = end + 1
		return
	}

	for {
		ref := s.getKey(index)
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
