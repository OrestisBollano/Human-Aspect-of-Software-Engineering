# 🚀 Go Coding Experiment

Welcome, participants! This repository contains a few small coding tasks designed to be completed in a GitHub Codespace.

Your goal is to follow the instructions, complete the tasks, and make sure all the automated tests pass.

## Getting Started: Your Personal Workspace

Please follow these steps to create your own private, separate coding environment.

1.  **Accept the Invitation:** You should have received an email from GitHub. Please **accept the collaborator invitation** first.
2.  **Go to the Repository:** Navigate to the main page for this project.
3.  **Create Your Codespace:**
    * Click the green **`< > Code`** button.
    * Select the **"Codespaces"** tab.
    * Click **"Create codespace on main"**.
4.  **Wait for Build:** GitHub will now build your environment. It will have Go installed and be ready for you in 1-2 minutes.

## Your Workflow: How to Complete a Task

For *each* task, you will follow this three-step process:

1.  **Navigate to the task folder** in your terminal.
    ```bash
    # Example for task 1
    cd task1
    ```
2.  **Run the test and see it fail.**
    ```bash
    go test
    ```
    This command runs the automated tests. You will see a `FAIL` message, which tells you what's wrong.
3.  **Code the solution.**
    * Open the `.go` file for that task (e.g., `task1/taskCheck.go`).
    * Write your code to solve the problem.
    * Save the file (`Ctrl+S` or `Cmd+S`).
4.  **Run the test again** (and again) until it passes.
    ```bash
    go test
    ```
    When your solution is correct, you will see a green `PASS` message.

> **Important:** Do not modify the `_test.go` files! Your goal is to make the existing tests pass by editing the *other* `.go` files (like `taskCheck.go` or `task2.go`).

---

## The Tasks

### Task 1: Python to Go Translation

1.  Navigate to the `task1` folder: `cd task1`
2.  Run `go test` to see the failing tests.
3.  Open and read the file `task1/README_TASK.md`. This contains a simple Python function.
4.  Your job is to translate that Python logic into Go. Write your Go code in the file `task1/taskCheck.go`.
5.  Run `go test` repeatedly until you see `PASS`.

### Task 2: Debugging

1.  Navigate to the `task2` folder: `cd ../task2`
2.  Run `go test` to see the failing tests.
3.  The file `task2/task2.go` is already implemented, but it **contains a bug**.
4.  Your job is to read the code in `task2/task2.go`, find the bug, and fix it.
5.  Run `go test` repeatedly until you see `PASS`.

---

## 🏁 Finishing the Experiment

When you have completed all tasks and all tests pass, please "submit" your work by saving it to the repository.

Run these three commands in your terminal:

```bash
git add .
git commit -m "Finished all tasks"
git push origin main
