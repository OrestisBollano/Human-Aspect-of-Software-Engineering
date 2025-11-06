### Task 1: Python to Go Translation

Your goal is to translate the following Python function into its Go equivalent in the `task1/taskCheck.go` file.

The Go function signature is already defined for you. You just need to implement the logic.

**Python Code to Translate:**
```python
def find_first(haystack: list[int], needle: int) -> int:
    """
    Finds the first index of 'needle' in 'haystack'.
    Returns -1 if not found.
    """
    for i, val in enumerate(haystack):
        if val == needle:
            return i
    return -1
