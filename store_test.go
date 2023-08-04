package bst

import "testing"

var (
	stringSink string
)

func TestStore_Set(t *testing.T) {
	type fields struct {
		kvs []KV[string, string]
	}

	type args struct {
		key   string
		value string
	}

	type testcase struct {
		name   string
		fields fields
		args   args
		want   []KV[string, string]
	}

	tests := []testcase{
		{
			name:   "basic",
			fields: fields{},
			args: args{
				key:   "0000",
				value: "1",
			},
			want: []KV[string, string]{
				{
					Key:   "0000",
					Value: "1",
				},
			},
		},
		{
			name: "overwrite",
			fields: fields{
				kvs: []KV[string, string]{
					{
						Key:   "0000",
						Value: "1",
					},
				},
			},
			args: args{
				key:   "0000",
				value: "2",
			},
			want: []KV[string, string]{
				{
					Key:   "0000",
					Value: "2",
				},
			},
		},
		{
			name: "prepend",
			fields: fields{
				kvs: []KV[string, string]{
					{
						Key:   "0001",
						Value: "1",
					},
				},
			},
			args: args{
				key:   "0000",
				value: "0",
			},
			want: []KV[string, string]{
				{
					Key:   "0000",
					Value: "0",
				},
				{
					Key:   "0001",
					Value: "1",
				},
			},
		},
		{
			name: "append",
			fields: fields{
				kvs: []KV[string, string]{
					{
						Key:   "0000",
						Value: "0",
					},
				},
			},
			args: args{
				key:   "0001",
				value: "1",
			},
			want: []KV[string, string]{
				{
					Key:   "0000",
					Value: "0",
				},
				{
					Key:   "0001",
					Value: "1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := New(tt.fields.kvs...)
			k.Set(tt.args.key, tt.args.value)
		})
	}
}

func Benchmark_Store_getIndex(b *testing.B) {
	var kvs []KV[string, string]
	for _, key := range testLetters {
		var kv KV[string, string]
		kv.Key = key
		kv.Value = key
		kvs = append(kvs, kv)
	}

	k := New(kvs...)
	b.ResetTimer()
	var match bool
	for i := 0; i < b.N; i++ {
		for _, key := range testLetters {
			indexSink, match = k.getIndex(key)
			if !match {
				b.Fatalf("received non-match for <%s>", key)
			}
		}
	}

	b.ReportAllocs()
}

func Benchmark_Store_Get(b *testing.B) {
	var kvs []KV[string, string]
	for _, key := range testLetters {
		var kv KV[string, string]
		kv.Key = key
		kv.Value = key
		kvs = append(kvs, kv)
	}

	k := New(kvs...)
	b.ResetTimer()

	var match bool
	for i := 0; i < b.N; i++ {
		for _, key := range testLetters {
			stringSink, match = k.Get(key)
			if !match {
				b.Fatalf("received non-match for <%s>", key)
			}
		}
	}

	b.ReportAllocs()
}

func Benchmark_Map_Get(b *testing.B) {
	m := make(map[string]string)
	for _, key := range testLetters {
		m[key] = key
	}

	b.ResetTimer()

	var match bool
	for i := 0; i < b.N; i++ {
		for _, key := range testLetters {
			stringSink, match = m[key]
			if !match {
				b.Fatalf("received non-match for <%s>", key)
			}
		}
	}

	b.ReportAllocs()
}
