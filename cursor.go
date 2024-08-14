package bst

import (
	"log"

	"github.com/itsmontoya/mappedslice"
)

type Cursor[K any, V any] struct {
	r *Raw[K, V]
	c mappedslice.Cursor[KV[K, V]]
}

func (c *Cursor[K, V]) Seek(seekTo K) (kv KV[K, V], ok bool) {
	index, _, err := c.r.getIndex(seekTo)
	if err != nil {
		log.Printf("error getting index for <%v>: %v", seekTo, err)
		return
	}

	return c.c.Seek(index)
}

func (c *Cursor[K, V]) Next() (kv KV[K, V], ok bool) {
	return c.c.Next()
}

func (c *Cursor[K, V]) Prev() (kv KV[K, V], ok bool) {
	return c.c.Prev()
}

func (c *Cursor[K, V]) Close() {
	c.r = nil
	c.c = nil
}
