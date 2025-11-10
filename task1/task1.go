package task1

import "strings"

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
	// Create an empty map for word counts
	counts := make(map[string]int)

	// Split the input string by whitespace (ignores multiple spaces, tabs, etc.)
	words := strings.Fields(s)

	// Count each word
	for _, w := range words {
		counts[w]++
	}

	return counts
}
