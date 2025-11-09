### Task 2: Python to Go Translation (Count Letters)

Your goal is to translate the following Python function into its Go equivalent in the `task2/task2.go` file.

The Go function signature is already defined for you. You just need to implement the logic.

**Python Code to Translate:**
```python
# Define a function 'count_letters' that takes one argument, 's'
def count_letters(s):
    
    # Create an empty dictionary
    counts = {}
    
    # Loop over each character 'ch' in the lowercase version of the string 's'
    for ch in s.lower():
        
        # Check if the character is an alphabetic letter
        if ch.isalpha():
            
            # If it is, get its current count (defaulting to 0)
            # and increment it by 1
            counts[ch] = counts.get(ch, 0) + 1
            
    # Return the final dictionary of counts
    return counts
