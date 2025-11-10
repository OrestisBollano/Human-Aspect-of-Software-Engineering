package task2

import (
	"strings"
	"unicode"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {

	// Create an empty map (dictionary equivalent)
	counts := make(map[rune]int)

	// Loop over each character 'ch' in the lowercase version of the string 's'
	for _, ch := range strings.ToLower(s) {

		// Check if the character is an alphabetic letter
		if unicode.IsLetter(ch) {

			// If it is, get its current count (default 0) and increment it by 1
			counts[ch] = counts[ch] + 1
			// Or shorthand: counts[ch]++
		}
	}

	// Return the final map of counts
	return counts
}
