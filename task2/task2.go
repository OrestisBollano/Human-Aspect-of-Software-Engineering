package task2

import (
	"unicode"
)

// CountLetters takes a string and returns a map of each letter
// (as a rune) to its count, ignoring case and non-letter characters.
func CountLetters(s string) map[rune]int {
	// TODO: Implement the Go version of the Python function.


	letterCounts := make(map[rune]int)
	

	for _, char := range s {
		
		// 3. Test if the character is an alphabetic letter
		if unicode.IsLetter(char) {
			
			// 4. Convert the letter to lowercase for case-insensitive counting
			lowerChar := unicode.ToLower(char)

			// 5. Increment the count in the map
			letterCounts[lowerChar]++
		}
	}	

	return letterCounts // Replace this
}
