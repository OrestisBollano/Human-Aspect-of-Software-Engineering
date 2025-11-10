package task1

	
import (
	"strings"
)

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
	// TODO: Implement the Go version of the Python function.
	m := make(map[string]int)
	if (len(s)<1) {
        return m
    }
	var s_arr = strings.Split(s, " ")
	for j := 0; j < len(s_arr); j++ {
    	var word = s_arr[j]
		m[word] +=1
    }
	return m // Replace this
}