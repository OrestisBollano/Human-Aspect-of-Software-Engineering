package task1

import "strings"

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
	// TODO: Implement the Go version of the Python function.
	// Create an empty map to hold word counts
	counts := make(map[string]int)

	// Split the input string into words by spaces
	words := strings.Fields(s)

	// Loop through each word
	for _, w := range words {
		// Increment the count for each word
		counts[w]++
	}

	// Return the completed map
	return counts
}

