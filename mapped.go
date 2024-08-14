package bst

import (
	"github.com/itsmontoya/mappedslice"
)

func NewMapped[K any, V any](filepath string, compare func(K, K) int, kvs ...KV[K, V]) (out *Mapped[K, V], err error) {
	var m Mapped[K, V]
	if m.s, err = mappedslice.New[KV[K, V]](filepath); err != nil {
		return
	}

	m.Raw = makeRaw(compare, m.s, kvs)
	out = &m
	return
}

type Mapped[K any, V any] struct {
	s *mappedslice.Slice[KV[K, V]]
	Raw[K, V]
}

func (m *Mapped[K, V]) Close() (err error) {
	return m.s.Close()
}
