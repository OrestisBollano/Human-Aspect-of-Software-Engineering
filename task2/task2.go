package task2

import (
	"fmt"
	"strings"
	"unicode"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	// TODO: Implement the Go version of the Python function.
	m := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		m[s.ToLower[i]] +=1
    }
	return m 
}
