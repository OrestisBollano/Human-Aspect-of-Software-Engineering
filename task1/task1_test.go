package task1

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		input string
		want  map[string]int
	}{
		{"hello world hello", map[string]int{"hello": 2, "world": 1}},
		{"a b c a b", map[string]int{"a": 2, "b": 2, "c": 1}},
		{"", map[string]int{}},
		{"one", map[string]int{"one": 1}},
		{"go go go", map[string]int{"go": 3}},
	}

	for _, tc := range cases {
		got := WordCount(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("WordCount(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
