package bst

import (
	"sort"
)

// NewStore a new Store instance
func NewStore[T any](kvs ...KV[T]) *Store[T] {
	s := makeStore[T](kvs)
	return &s
}

func makeStore[T any](kvs []KV[T]) (s Store[T]) {
	sz := len(kvs)
	if sz == 0 {
		sz = 8
	}

	s.kvs = make([]KV[T], len(kvs), sz)
	copy(s.kvs, kvs)
	sort.Slice(s.kvs, func(i, j int) (less bool) {
		return s.kvs[i].Key < s.kvs[j].Key
	})

	return
}

// Store is a Store Stringset
type Store[T any] struct {
	kvs []KV[T]
}

// Set will place a key
func (s *Store[T]) Set(key string, value T) {
	index, match := s.getIndex(key)
	if match {
		s.kvs[index].Value = value
		return
	}

	pair := makeKV(key, value)
	first := s.kvs[:index]
	second := append([]KV[T]{pair}, s.kvs[index:]...)
	s.kvs = append(first, second...)
}

// Get will retrieve a value for a given key
func (s *Store[T]) Get(key string) (value T, has bool) {
	var index int
	if index, has = s.getIndex(key); !has {
		return
	}

	value = s.kvs[index].Value
	return
}

// UsSet will remove a key
func (s *Store[T]) Unset(key string) {
	index, match := s.getIndex(key)
	if !match {
		return
	}

	first := s.kvs[:index]
	second := s.kvs[index+1:]
	s.kvs = append(first, second...)
}

// Has will return if a key exists
func (s *Store[T]) Has(key string) (has bool) {
	_, has = s.getIndex(key)
	return
}

// Len will return the keys length
func (s *Store[T]) Len() (n int) {
	return len(s.kvs)
}

// Len will return the keys length
func (s *Store[T]) Slice() (kvs []KV[T]) {
	kvs = make([]KV[T], len(s.kvs))
	copy(kvs, s.kvs)
	return
}

// Len will return the keys length
func (s *Store[T]) ForEach(fn func(string, interface{}) (end bool)) (ended bool) {
	for _, kv := range s.kvs {
		if ended = fn(kv.Key, kv.Value); ended {
			return
		}
	}

	return
}

func (s *Store[T]) getKey(index int) string {
	return s.kvs[index].Key
}

func (s *Store[T]) getIndex(key string) (index int, match bool) {
	sz := s.Len()
	if sz == 0 {
		return
	}

	start := 0
	end := sz - 1
	index = sz / 2

	if s.getKey(end) < key {
		index = end + 1
		return
	}

	for {
		ref := s.getKey(index)
		switch {
		case key == ref:
			match = true
			return
		case start == end:
			index++
			return
		case key < ref:
			end = index - 1
		case key > ref:
			start = index + 1
		}

		index = (start + end) / 2
	}
}
