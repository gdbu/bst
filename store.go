package bst

import "sort"

// NewStore a new Store instance
func NewStore(kvs ...KV) *Store {
	s := makeStore(kvs)
	return &s
}

func makeStore(kvs []KV) (k Store) {
	sz := len(kvs)
	if sz == 0 {
		sz = 8
	}

	k.s = make([]KV, len(kvs), sz)
	copy(k.s, kvs)
	sort.Slice(k.s, func(i, j int) (less bool) {
		return k.s[i].Key < k.s[j].Key
	})

	return
}

// Store is a Store Stringset
type Store struct {
	s []KV
}

// Set will place a key
func (k *Store) Set(key string, value interface{}) {
	index, match := getIndex(k, key)
	if match {
		return
	}

	pair := makeKV(key, value)
	first := k.s[:index]
	second := append([]KV{pair}, k.s[index:]...)
	k.s = append(first, second...)
}

// Get will retrieve a value for a given key
func (k *Store) Get(key string) (value interface{}, has bool) {
	var index int
	if index, has = getIndex(k, key); !has {
		return
	}

	value = k.s[index].Value
	return
}

// UsSet will remove a key
func (k *Store) Unset(key string) {
	index, match := getIndex(k, key)
	if !match {
		return
	}

	first := k.s[:index]
	second := k.s[index+1:]
	k.s = append(first, second...)
}

// Has will return if a key exists
func (k *Store) Has(key string) (has bool) {
	_, has = getIndex(k, key)
	return
}

// Len will return the keys length
func (k *Store) Len() (n int) {
	return len(k.s)
}

// Len will return the keys length
func (k *Store) Slice() (s []KV) {
	s = make([]KV, len(k.s))
	copy(s, k.s)
	return
}

// Len will return the keys length
func (k *Store) ForEach(fn func(string, interface{}) (end bool)) (ended bool) {
	for _, kv := range k.s {
		if ended = fn(kv.Key, kv.Value); ended {
			return
		}
	}

	return
}

func (k *Store) getKey(index int) string {
	return k.s[index].Key
}
