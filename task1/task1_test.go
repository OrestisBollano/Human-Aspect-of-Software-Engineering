package task1

import "testing"

func TestFindFirst(t *testing.T) {
	cases := []struct {
		haystack []int
		needle   int
		want     int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{10, 20, 30}, 5, -1},
		{[]int{5, 5, 5, 5}, 5, 0},
		{[]int{}, 1, -1},
		{[]int{1, 2, 3}, 1, 0},
		{[]int{1, 2, 3}, 5, -1},
	}

	for _, tc := range cases {
		got := FindFirst(tc.haystack, tc.needle)
		if got != tc.want {
			t.Errorf("FindFirst(%v, %d) = %d, want %d", tc.haystack, tc.needle, got, tc.want)
		}
	}
}
