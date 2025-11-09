### Task 2: Debugging (Reverse String)

Your goal for this task is to find and fix a bug in an existing Go function.

#### The Problem

The file `task2/task2.go` contains a function `Reverse(s string)` that is *supposed* to take a string and return that string in reverse order.

For example:
* `Reverse("hello")` should return `"olleh"`
* `Reverse("Go")` should return `"oG"`

However, the function has a bug and is failing some of the tests.

#### Your Job

1.  Run `go test` in the `task2` terminal to see which tests are failing.
2.  Read the code in `task2/task2.go`.
3.  Find the logical error in the code.
4.  Fix the bug directly in the `task2.go` file.
5.  Run `go test` again. If it passes, you are done with this task!
