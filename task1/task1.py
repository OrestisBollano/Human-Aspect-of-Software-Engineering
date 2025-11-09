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

# This part only runs when you execute this file directly
if __name__ == "__main__":
    text = "hello world hello"
    print(f"Input: '{text}'")
    print("Output:", word_counts(text))
    # To run this, type 'python3 task1.py' in your terminal
