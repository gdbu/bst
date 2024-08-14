package bst

import (
	"os"
	"reflect"
	"slices"
	"testing"
)

func TestMapped_Set(t *testing.T) {
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

func TestMapped_Update(t *testing.T) {
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

		want         []KV[keytype, int]
		wantExisting int
		wantErr      bool
	}

	tests := []testcase{
		{
			name: "basic",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 0,
					},
				},
			},
			args: args{
				key:   keytype{0, 0, 0, 0},
				value: 1,
			},
			want: []KV[keytype, int]{
				makeKV(keytype{0, 0, 0, 0}, 1),
			},
			wantExisting: 0,
		},

		{
			name: "multiple",
			fields: fields{
				kvs: []KV[keytype, int]{
					makeKV(keytype{0, 0, 0, 0}, 0),
					makeKV(keytype{0, 0, 0, 1}, 1),
					makeKV(keytype{0, 0, 0, 2}, 2),
					makeKV(keytype{0, 0, 0, 3}, 3),
					makeKV(keytype{0, 0, 0, 4}, 4),
					makeKV(keytype{0, 0, 0, 5}, 5),
					makeKV(keytype{0, 0, 0, 6}, 6),
					makeKV(keytype{0, 0, 0, 7}, 7),
					makeKV(keytype{0, 0, 0, 8}, 8),
					makeKV(keytype{0, 0, 0, 9}, 9),
				},
			},
			args: args{
				key:   keytype{0, 0, 0, 9},
				value: 1337,
			},
			want: []KV[keytype, int]{
				makeKV(keytype{0, 0, 0, 0}, 0),
				makeKV(keytype{0, 0, 0, 1}, 1),
				makeKV(keytype{0, 0, 0, 2}, 2),
				makeKV(keytype{0, 0, 0, 3}, 3),
				makeKV(keytype{0, 0, 0, 4}, 4),
				makeKV(keytype{0, 0, 0, 5}, 5),
				makeKV(keytype{0, 0, 0, 6}, 6),
				makeKV(keytype{0, 0, 0, 7}, 7),
				makeKV(keytype{0, 0, 0, 8}, 8),
				makeKV(keytype{0, 0, 0, 9}, 1337),
			},
			wantExisting: 9,
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

			var gotExisting int
			err = m.Update(tt.args.key, func(existing int) (out int) {
				gotExisting = existing
				return tt.args.value
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("Mapped.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotExisting != tt.wantExisting {
				t.Errorf("Mapped.Update() gotExisting = %v, wantExisting %v", gotExisting, tt.wantExisting)
				return
			}
			got := m.b.Slice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mapped.Update() got = %v, want %v", got, tt.want)
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
				t.Errorf("Mapped.Update() (after close and re-open) got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapped_Get(t *testing.T) {
	type keytype [4]int
	type fields struct {
		kvs []KV[keytype, int]
	}

	type args struct {
		key keytype
	}

	type testcase struct {
		name   string
		fields fields
		args   args

		want    int
		wantErr bool
	}

	tests := []testcase{
		{
			name: "get first",
			fields: fields{
				kvs: []KV[keytype, int]{
					makeKV(keytype{0, 0, 0, 0}, 0),
					makeKV(keytype{0, 0, 0, 1}, 1),
					makeKV(keytype{0, 0, 0, 2}, 2),
					makeKV(keytype{0, 0, 0, 3}, 3),
					makeKV(keytype{0, 0, 0, 4}, 4),
					makeKV(keytype{0, 0, 0, 5}, 5),
					makeKV(keytype{0, 0, 0, 6}, 6),
					makeKV(keytype{0, 0, 0, 7}, 7),
					makeKV(keytype{0, 0, 0, 8}, 8),
					makeKV(keytype{0, 0, 0, 9}, 9),
				},
			},
			args: args{
				key: keytype{0, 0, 0, 0},
			},
			want: 0,
		},
		{
			name: "get middle",
			fields: fields{
				kvs: []KV[keytype, int]{
					makeKV(keytype{0, 0, 0, 0}, 0),
					makeKV(keytype{0, 0, 0, 1}, 1),
					makeKV(keytype{0, 0, 0, 2}, 2),
					makeKV(keytype{0, 0, 0, 3}, 3),
					makeKV(keytype{0, 0, 0, 4}, 4),
					makeKV(keytype{0, 0, 0, 5}, 5),
					makeKV(keytype{0, 0, 0, 6}, 6),
					makeKV(keytype{0, 0, 0, 7}, 7),
					makeKV(keytype{0, 0, 0, 8}, 8),
					makeKV(keytype{0, 0, 0, 9}, 9),
				},
			},
			args: args{
				key: keytype{0, 0, 0, 5},
			},
			want: 5,
		},
		{
			name: "get last",
			fields: fields{
				kvs: []KV[keytype, int]{
					makeKV(keytype{0, 0, 0, 0}, 0),
					makeKV(keytype{0, 0, 0, 1}, 1),
					makeKV(keytype{0, 0, 0, 2}, 2),
					makeKV(keytype{0, 0, 0, 3}, 3),
					makeKV(keytype{0, 0, 0, 4}, 4),
					makeKV(keytype{0, 0, 0, 5}, 5),
					makeKV(keytype{0, 0, 0, 6}, 6),
					makeKV(keytype{0, 0, 0, 7}, 7),
					makeKV(keytype{0, 0, 0, 8}, 8),
					makeKV(keytype{0, 0, 0, 9}, 9),
				},
			},
			args: args{
				key: keytype{0, 0, 0, 9},
			},
			want: 9,
		},
		{
			name: "get non-existing key",
			fields: fields{
				kvs: []KV[keytype, int]{
					makeKV(keytype{0, 0, 0, 0}, 0),
					makeKV(keytype{0, 0, 0, 1}, 1),
					makeKV(keytype{0, 0, 0, 2}, 2),
					makeKV(keytype{0, 0, 0, 3}, 3),
					makeKV(keytype{0, 0, 0, 4}, 4),
					makeKV(keytype{0, 0, 0, 5}, 5),
					makeKV(keytype{0, 0, 0, 6}, 6),
					makeKV(keytype{0, 0, 0, 7}, 7),
					makeKV(keytype{0, 0, 0, 8}, 8),
					makeKV(keytype{0, 0, 0, 9}, 9),
				},
			},
			args: args{
				key: keytype{0, 0, 0, 11},
			},
			want:    0,
			wantErr: true,
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

			var got int
			got, err = m.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapped.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Mapped.Get() got = %v, want %v", got, tt.want)
			}

			if err = m.Close(); err != nil {
				t.Error(err)
			}

			if m, err = NewMapped[keytype, int]("./test.bat", func(a, b keytype) (compare int) {
				return slices.Compare(a[:], b[:])
			}); err != nil {
				t.Error(err)
			}

			got, err = m.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapped.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Mapped.Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapped_RemoveAt(t *testing.T) {
	type keytype [4]int
	type fields struct {
		kvs []KV[keytype, int]
	}

	type args struct {
		key keytype
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
			name: "basic",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 1,
					},
				},
			},
			args: args{
				key: keytype{0, 0, 0, 0},
			},
			want: []KV[keytype, int]{},
		},
		{
			name: "remove first",
			fields: fields{
				kvs: []KV[keytype, int]{
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
			args: args{
				key: keytype{0, 0, 0, 0},
			},
			want: []KV[keytype, int]{
				{
					Key:   keytype{0, 0, 0, 1},
					Value: 1,
				},
			},
		},
		{
			name: "remove last",
			fields: fields{
				kvs: []KV[keytype, int]{
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
			args: args{
				key: keytype{0, 0, 0, 1},
			},
			want: []KV[keytype, int]{
				{
					Key:   keytype{0, 0, 0, 0},
					Value: 0,
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

			err = m.RemoveAt(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapped.RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := m.b.Slice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mapped.RemoveAt() got = %v, want %v", got, tt.want)
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

func TestMapped_ForEach(t *testing.T) {
	type keytype [4]int
	type fields struct {
		kvs []KV[keytype, int]
		end bool
	}

	type testcase struct {
		name   string
		fields fields

		wantKeys  []keytype
		wantEnded bool
	}

	tests := []testcase{
		{
			name: "single",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 0,
					},
				},
			},
			wantKeys: []keytype{
				{0, 0, 0, 0},
			},
		},
		{
			name: "multiple",
			fields: fields{
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 0,
					},
					{
						Key:   keytype{0, 0, 0, 1},
						Value: 1,
					},
					{
						Key:   keytype{0, 0, 0, 2},
						Value: 2,
					},
					{
						Key:   keytype{0, 0, 0, 3},
						Value: 3,
					},
					{
						Key:   keytype{0, 0, 0, 4},
						Value: 4,
					},
					{
						Key:   keytype{0, 0, 0, 5},
						Value: 5,
					},
					{
						Key:   keytype{0, 0, 0, 6},
						Value: 6,
					},
					{
						Key:   keytype{0, 0, 0, 7},
						Value: 7,
					},
				},
			},
			wantKeys: []keytype{
				{0, 0, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 2},
				{0, 0, 0, 3},
				{0, 0, 0, 4},
				{0, 0, 0, 5},
				{0, 0, 0, 6},
				{0, 0, 0, 7},
			},
		},
		{
			name: "multiple with end",
			fields: fields{
				end: true,
				kvs: []KV[keytype, int]{
					{
						Key:   keytype{0, 0, 0, 0},
						Value: 0,
					},
					{
						Key:   keytype{0, 0, 0, 1},
						Value: 1,
					},
					{
						Key:   keytype{0, 0, 0, 2},
						Value: 2,
					},
					{
						Key:   keytype{0, 0, 0, 3},
						Value: 3,
					},
					{
						Key:   keytype{0, 0, 0, 4},
						Value: 4,
					},
					{
						Key:   keytype{0, 0, 0, 5},
						Value: 5,
					},
					{
						Key:   keytype{0, 0, 0, 6},
						Value: 6,
					},
					{
						Key:   keytype{0, 0, 0, 7},
						Value: 7,
					},
				},
			},
			wantKeys: []keytype{
				{0, 0, 0, 0},
			},
			wantEnded: true,
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

			var gotKeys []keytype
			gotEnded := m.ForEach(func(k keytype, v int) (end bool) {
				gotKeys = append(gotKeys, k)
				return tt.fields.end
			})

			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("Mapped.ForEach() gotKeys = %v, wantKeys %v", gotKeys, tt.wantKeys)
				return
			}

			if gotEnded != tt.wantEnded {
				t.Errorf("Mapped.ForEach() gotKeys = %v, wantKeys %v", gotKeys, tt.wantEnded)
				return
			}

			if err = m.Close(); err != nil {
				t.Error(err)
				return
			}

			if m, err = NewMapped[keytype, int]("./test.bat", func(a, b keytype) (compare int) {
				return slices.Compare(a[:], b[:])
			}); err != nil {
				t.Error(err)
			}

			gotKeys = gotKeys[:0]
			m.ForEach(func(k keytype, v int) (end bool) {
				gotKeys = append(gotKeys, k)
				return tt.fields.end
			})

			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("Mapped.ForEach() (after close and re-open) gotKeys = %v, wantKeys %v", gotKeys, tt.wantKeys)
			}
		})
	}
}
