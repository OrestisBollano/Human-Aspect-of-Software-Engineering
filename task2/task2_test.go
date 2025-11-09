package task2

import (
	"reflect"
	"testing"
)

func TestCountLetters(t *testing.T) {
	cases := []struct {
		input string
		want  map[rune]int
	}{
		{
			"Hello World",
			map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, 'w': 1, 'r': 1, 'd': 1},
		},
		{
			"aA bB 123!",
			map[rune]int{'a': 2, 'b': 2},
		},
		{
			"",
			map[rune]int{},
		},
		{
			"Gophers",
			map[rune]int{'g': 1, 'o': 1, 'p': 1, 'h': 1, 'e': 1, 'r': 1, 's': 1},
		},
	}

	for _, tc := range cases {
		got := CountLetters(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("CountLetters(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
