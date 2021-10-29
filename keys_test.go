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
		has := k.Has(tc.key)
		if tc.wantMatch != has {
			t.Fatalf("invalid match, expected %v and received %v", tc.wantMatch, has)
		}

		if !has {
			k.Set("foo")
		}

		if !k.Has("foo") {
			t.Fatal("does not have foo when expected")
		}

		k.Unset("foo")

		if k.Has("foo") {
			t.Fatal("has foo when not expected")
		}
	}

}

func TestKeys_getIndex(t *testing.T) {
	var k Keys
	type testcase struct {
		strs      sort.StringSlice
		key       string
		wantIndex int
		wantMatch bool
	}

	appendTests := func(strs []string) (tcs []testcase) {
		for i, str := range strs {
			var tc testcase
			tc.strs = strs
			tc.key = str
			tc.wantIndex = i
			tc.wantMatch = true
		}

		return
	}

	tcs := []testcase{
		{
			strs:      testLetters,
			key:       "foo",
			wantIndex: 6,
			wantMatch: false,
		},
		{
			strs:      testLetters,
			key:       "zoinks",
			wantIndex: 26,
			wantMatch: false,
		},
	}

	tcs = append(tcs, appendTests(testLetters)...)

	for i, tc := range tcs {
		sort.Sort(tc.strs)
		k.s = tc.strs

		index, match := k.getIndex(tc.key)
		switch {
		case tc.wantIndex != index:
			t.Fatalf("invalid index, expected %d and received %d (test case index %d)", tc.wantIndex, index, i)
		case tc.wantMatch != match:
			t.Fatalf("invalid match, expected %v and received %v (test case index %d)", tc.wantMatch, match, i)
		}
	}
}
