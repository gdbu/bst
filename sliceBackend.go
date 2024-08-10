package bst

var _ Backend[int, int] = &sliceBackend[int, int]{}

type sliceBackend[K, V any] []KV[K, V]

func (s *sliceBackend[K, V]) Get(index int) (kv KV[K, V]) {
	ref := *s
	kv = ref[index]
	return
}

func (s *sliceBackend[K, V]) Set(index int, v V) {
	ref := *s
	ref[index].Value = v
}

func (s *sliceBackend[K, V]) Unset(index int) {
	ref := *s
	first := ref[:index]
	second := ref[index+1:]
	ref = append(first, second...)
	*s = ref
}

func (s *sliceBackend[K, V]) InsertAt(index int, pair KV[K, V]) {
	ref := *s
	first := ref[:index]
	second := append([]KV[K, V]{pair}, ref[index:]...)
	ref = append(first, second...)
	*s = ref
}

func (s *sliceBackend[K, V]) Len() int {
	return len(*s)
}

func (s *sliceBackend[K, V]) Slice() (out []KV[K, V]) {
	ref := *s
	out = make([]KV[K, V], len(ref))
	copy(out, ref)
	return
}

func (s *sliceBackend[K, V]) ForEach(fn func(K, V) (end bool)) (ended bool) {
	ref := *s
	for _, kv := range ref {
		if ended = fn(kv.Key, kv.Value); ended {
			return
		}
	}

	return
}
