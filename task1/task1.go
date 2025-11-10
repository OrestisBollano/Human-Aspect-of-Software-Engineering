package task1

import "strings"

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
	// TODO: Implement the Go version of the Python function.
	wordCounts := make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		// The key operation: Increment the count for the current word.
		// If the word is seen for the first time, Go initializes the
		// integer value to its zero value (0), so the first increment
		// correctly sets the count to 1.
		wordCounts[word]++
	}	
	return wordCounts // Replace this
}
