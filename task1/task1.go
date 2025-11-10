package task1

import "strings"

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
		// TODO: Implement the Go version of the Python function.
	
	m := make(map[string]int)
	
	for w := string.SplitN(s, " ") {
		v1 := m[w]
		m[w] := v1 + 1
	}

	return m // Replace this
}