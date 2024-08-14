package bst

import (
	"os"
	"reflect"
	"slices"
	"testing"
)

func TestMapped_Set_(t *testing.T) {
	type keytype [4]int
	type fields struct {
		kvs []KV[keytype, int]
	}

	type args struct {
		key   keytype
		value int
	}

	type testcase struct {
		name   string
		fields fields
		args   args

		want    []KV[keytype, int]
		wantErr bool
	}

	tests := []testcase{
		{
			name:   "basic",
			fields: fields{},
			args: args{
				key:   keytype{0, 0, 0, 0},
				value: 1,
			},
			want: []KV[keytype, int]{
				{
					Key:   keytype{0, 0, 0, 0},
					Value: 1,
				},
			},
		},
		{
			name: "overwrite",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 1,
					},
				},
			},
			args: args{
				key:   keytype{0, 0, 0, 0},
				value: 2,
			},
			want: []KV[keytype, int]{
				{
					Key:   keytype{0, 0, 0, 0},
					Value: 2,
				},
			},
		},
		{
			name: "prepend",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 1},
						Value: 1,
					},
				},
			},
			args: args{
				key:   keytype{0, 0, 0, 0},
				value: 0,
			},
			want: []KV[keytype, int]{
				{
					Key:   keytype{0, 0, 0, 0},
					Value: 0,
				},
				{
					Key:   keytype{0, 0, 0, 1},
					Value: 1,
				},
			},
		},
		{
			name: "append",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 0,
					},
				},
			},
			args: args{
				key:   keytype{0, 0, 0, 1},
				value: 1,
			},
			want: []KV[keytype, int]{
				{
					Key:   keytype{0, 0, 0, 0},
					Value: 0,
				},
				{
					Key:   keytype{0, 0, 0, 1},
					Value: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				m   *Mapped[keytype, int]
				err error
			)

			if m, err = NewMapped("./test.bat", func(a, b keytype) (compare int) {
				return slices.Compare(a[:], b[:])
			}, tt.fields.kvs...); err != nil {
				t.Error(err)
				return
			}
			defer os.Remove("./test.bat")

			err = m.Set(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapped.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := m.b.Slice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mapped.Set() got = %v, want %v", got, tt.want)
			}

			if err = m.Close(); err != nil {
				t.Error(err)
			}

			if m, err = NewMapped[keytype, int]("./test.bat", func(a, b keytype) (compare int) {
				return slices.Compare(a[:], b[:])
			}); err != nil {
				t.Error(err)
			}

			got = m.b.Slice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mapped.Set() (after close and re-open) got = %v, want %v", got, tt.want)
			}
		})
	}
}
