package bst

import (
	"sort"
	"testing"
)

var (
	testLetters = []string{
		"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z",
	}

	indexSink int
)

func TestKeys(t *testing.T) {
	type testcase struct {
		strs      sort.StringSlice
		key       string
		wantMatch bool
	}

	appendTests := func(strs []string) (tcs []testcase) {
		for _, str := range strs {
			var tc testcase
			tc.strs = strs
			tc.key = str
			tc.wantMatch = true
			tcs = append(tcs, tc)
		}

		return
	}

	tcs := []testcase{
		{
			strs:      testLetters,
			key:       "foo",
			wantMatch: false,
		},
	}

	tcs = append(tcs, appendTests(testLetters)...)

	for _, tc := range tcs {
		k := NewKeys(tc.strs...)
		has, err := k.Has(tc.key)
		if err != nil {
			t.Fatal(err)
		}

		if tc.wantMatch != has {
			t.Fatalf("invalid match, expected %v and received %v", tc.wantMatch, has)
		}

		if !has {
			k.Set("foo")
			has, err := k.Has("foo")
			if err != nil {
				t.Fatal(err)
			}

			if !has {
				t.Fatal("does not have foo when expected")
			}

			k.Unset("foo")

		}

		has, err = k.Has("foo")
		if err != nil {
			t.Fatal(err)
		}

		if has {
			t.Fatal("has foo when not expected")
		}
	}

}
