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

# This part only runs when you execute this file directly
if __name__ == "__main__":
    text = "Hello World"
    print(f"Input: '{text}'")
    print("Output:", count_letters(text))
    # To run this, type 'python3 task2.py' in your terminal
