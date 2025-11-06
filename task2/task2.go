package task2

// Reverse takes a string and returns it in reverse order.
// NOTE: This function contains a bug!
func Reverse(s string) string {
	runes := []rune(s)
	// The bug is in the loop condition.
	// It stops one character too early, so it fails
	// to swap the middle letters on odd-length strings
	// and fails on short strings.
	for i, j := 0, len(runes)-1; i < j-1; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
