package bst

import "testing"

func Benchmark_getIndex(b *testing.B) {
	k := makeKeys(testLetters)
	b.ResetTimer()
	var match bool
	for i := 0; i < b.N; i++ {
		for _, key := range testLetters {
			indexSink, match = getIndex(&k, key)
			if !match {
				b.Fatalf("received non-match for <%s>", key)
			}
		}
	}

	b.ReportAllocs()
}
