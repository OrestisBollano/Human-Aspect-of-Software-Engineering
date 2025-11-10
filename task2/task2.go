package task2

import (
	"strings"
	"unicode"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {
	// TODO: Implement the Go version of the Python function.
	
	// Create an empty map to store letter counts
	counts := make(map[rune]int)

	// Convert the string to lowercase
	lower := strings.ToLower(s)

	// Loop over each rune (character) in the string
	for _, ch := range lower {
		// Check if the character is a letter
		if unicode.IsLetter(ch) {
			// Increment the count for this letter
			counts[ch]++
		}
	}

	// Return the completed map
	return counts
}
