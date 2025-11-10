package task2

import (
	"strings"
	"unicode"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {
    counts := make(map[rune]int)

    for _, ch := range strings.ToLower(s) {
        if unicode.IsLetter(ch) {
            counts[ch]++
        }
    }

    return counts
}
