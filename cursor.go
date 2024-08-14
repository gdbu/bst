package bst

import "log"

type Cursor[K any, V any] struct {
	r *Raw[K, V]
	c BackendCursor[KV[K, V]]
}

func (c *Cursor[K, V]) Seek(k K) (v V, ok bool) {
	index, _, err := c.r.getIndex(k)
	if err != nil {
		log.Printf("error getting index for <%v>: %v", k, err)
	}

	var kv KV[K, V]
	if kv, ok = c.c.Seek(index); !ok {
		return
	}

	v = kv.Value
	return
}

func (c *Cursor[K, V]) Next() (v V, ok bool) {
	var kv KV[K, V]
	if kv, ok = c.c.Next(); !ok {
		return
	}

	v = kv.Value
	return
}

func (c *Cursor[K, V]) Prev() (v V, ok bool) {
	var kv KV[K, V]
	if kv, ok = c.c.Prev(); !ok {
		return
	}

	v = kv.Value
	return
}

func (c *Cursor[K, V]) Close() {
	c.r = nil
	c.c = nil
}
