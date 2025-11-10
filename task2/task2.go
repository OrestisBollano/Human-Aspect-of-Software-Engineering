package task2

import (
	"unicode"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {
	// Create an empty map to store letter counts
	counts := make(map[rune]int)

	// Loop over each character in lowercase form
	for _, ch := range s {
		// Convert to lowercase
		ch = unicode.ToLower(ch)

		// Check if it's an alphabetic letter
		if unicode.IsLetter(ch) {
			counts[ch]++
		}
	}

	return counts
}
