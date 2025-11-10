package task1

import "strings"

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
	// TODO: Implement the Go version of the Python function.
	counts := make(map[string]int)

	for _, w := range strings.Fields(s) {
        counts[w]++
    }
	
	return counts // Replace this
}
