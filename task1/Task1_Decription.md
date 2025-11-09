### Task 1: Python to Go Translation (Word Counter)

Your goal is to translate the following Python function into its Go equivalent in the `task1/task1.go` file.

The Go function signature is already defined for you. You just need to implement the logic.

**Python Code to Translate:**
```python
# Define a function named 'word_counts' that takes one argument, 'text'
def word_counts(text):
    
    # Create an empty dictionary (in Go, this will be a 'map')
    counts = {}
    
    # Split the input 'text' into a list of words (by spaces)
    # and loop through each word 'w' in that list
    for w in text.split():
        
        # For each word 'w', try to get its current count from the 'counts' dictionary.
        # If the word isn't in the dictionary yet, '.get(w, 0)' returns a default value of 0.
        # Then, add 1 to that count and assign it back to the dictionary.
        counts[w] = counts.get(w, 0) + 1
        
    # After the loop is finished, return the completed 'counts' dictionary
    return counts
