package task2

import (
	"regexp"
	"strings"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {
	// TODO: Implement the Go version of the Python function.
	m := make(map[rune]int)
	re := regexp.MustCompile(`[^a-z]`)
	for _, letter := range strings.ToLower(s) {
		if !re.MatchString(string(letter)) {
			m[letter]++
		}
	}
	return m
}
