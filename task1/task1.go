package task1

import  "strings"
	

// WordCount takes a string, splits it into words,
// and returns a map of each word to its count.
func WordCount(s string) map[string]int {
	// TODO: Implement the Go version of the Python function.


    counts := make(map[string]int)
	    b := [...]string{strings.Split(s, " ")}
	
	for w := range b {
		if(counts[b[w]]){
			counts[b[w]] = counts[b[w]] + 1
		}else{
			counts[b[w]] = 0
		}

	}
	return counts // Replace this
	
}
