package bst

import "testing"

func TestStore_Set(t *testing.T) {
	type fields struct {
		s []KV
	}

	type args struct {
		key   string
		value interface{}
	}

	type testcase struct {
		name   string
		fields fields
		args   args
		want   []KV
	}

	tests := []testcase{
		{
			name:   "basic",
			fields: fields{},
			args: args{
				key:   "0000",
				value: "1",
			},
			want: []KV{
				{
					Key:   "0000",
					Value: "1",
				},
			},
		},
		{
			name: "overwrite",
			fields: fields{
				s: []KV{
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
			want: []KV{
				{
					Key:   "0000",
					Value: "2",
				},
			},
		},
		{
			name: "prepend",
			fields: fields{
				s: []KV{
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
			want: []KV{
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
				s: []KV{
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
			want: []KV{
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
			k := &Store{
				s: tt.fields.s,
			}
			k.Set(tt.args.key, tt.args.value)
		})
	}
}

func Benchmark_Store_getIndex(b *testing.B) {
	k := NewKeys(testLetters...)
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
