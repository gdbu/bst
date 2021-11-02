package bst

import "sort"

// NewKeys a new Keys instance
func NewKeys(keys ...string) *Keys {
	s := makeKeys(keys)
	return &s
}

func makeKeys(keys []string) (k Keys) {
	sz := len(keys)
	if sz == 0 {
		sz = 8
	}

	k.s = make([]string, len(keys), sz)
	copy(k.s, keys)
	sort.Strings(k.s)
	return
}

// Keys is a Keys Stringset
type Keys struct {
	s []string
}

// Set will place a key
func (k *Keys) Set(key string) {
	index, match := k.getIndex(key)
	if match {
		return
	}

	first := k.s[:index]
	second := append([]string{key}, k.s[index:]...)
	k.s = append(first, second...)
}

// UsSet will remove a key
func (k *Keys) Unset(key string) {
	index, match := k.getIndex(key)
	if !match {
		return
	}

	first := k.s[:index]
	second := k.s[index+1:]
	k.s = append(first, second...)
}

// Has will return if a key exists
func (k *Keys) Has(key string) (has bool) {
	_, has = k.getIndex(key)
	return
}

// Len will return the keys length
func (k *Keys) Len() (n int) {
	return len(k.s)
}

// Len will return the keys length
func (k *Keys) Slice() (s []string) {
	s = make([]string, len(k.s))
	copy(s, k.s)
	return
}

// Len will return the keys length
func (k *Keys) ForEach(fn func(string) (end bool)) (ended bool) {
	for _, k := range k.s {
		if ended = fn(k); ended {
			return
		}
	}

	return
}

func (k *Keys) getIndex(key string) (index int, match bool) {
	if len(k.s) == 0 {
		return
	}

	start := 0
	end := len(k.s) - 1
	index = len(k.s) / 2

	for {
		ref := k.s[index]
		switch {
		case key == ref:
			match = true
			return
		case key < ref:
			end = index
		case key > ref:
			start = index
		}

		switch {
		case start == end:
			if key > ref {
				index++
			}

			return
		case end-start > 1:
			index = (start + end) / 2
		case start == index:
			start++
			index++
		case end == index:
			end--
			index--
		}
	}
}

func (k *Keys) getKey(index int) string {
	return k.s[index]
}
