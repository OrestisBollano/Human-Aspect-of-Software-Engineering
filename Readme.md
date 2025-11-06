# 🚀 Go Coding Experiment

Welcome, participants! This repository contains a few small coding tasks designed to be completed in a GitHub Codespace.

Your goal is to follow the instructions for your group, complete the tasks, and make sure all the automated tests pass.

## ⚠️ Find Your Group Instructions First ⚠️

Please read the instructions for your assigned group **before** you begin.

---

### Group 1: Non-AI Group Rules

If you are in the **Non-AI Group**, you are **only** allowed to use the following websites to help you:

* The official Go Documentation: [golang.org/doc/](https://golang.org/doc/)
* Stack Overflow: [stackoverflow.com](https://stackoverflow.com)
* Go by Example: [gobyexample.com](https://gobyexample.com)

Please **do not** use any AI assistants, chatbots (like ChatGPT), or AI-powered search (like Perplexity).

---

### Group 2: AI Group Rules

If you are in the **AI Group**, you are encouraged to use any online resources to help you, including:

* The official Go Documentation: [golang.org/doc/](https://golang.org/doc/)
* Stack Overflow: [stackoverflow.com](https://stackoverflow.com)
* Go by Example: [gobyexample.com](https://gobyexample.com)
* **AI Chatbots:** You are **required** to use an online AI chatbot (like ChatGPT, Gemini, Copilot, etc.) to help you with the tasks.

#### **CRITICAL INSTRUCTIONS (AI Group Only)**

We need to collect the logs of your interaction with the AI chatbot.

1.  As you work, please use the AI chatbot for help on both tasks.
2.  When you are finished with the experiment, **you must download or export your full chat history** (e.g., in ChatGPT, go to Settings > Data Controls > Export data).
3.  After finishing, please **upload your downloaded log file** to this Microsoft Form:

    ➡️ **[INSERT YOUR MICROSOFT FORM LINK HERE]** ⬅️

This is a required step for the experiment.

---

## Getting Started: Your Personal Workspace (All Groups)

Please follow these steps to create your own private, separate coding environment.

1.  **Accept the Invitation:** You should have received an email from GitHub. Please **accept the collaborator invitation** first.
2.  **Go to the Repository:** Navigate to the main page for this project.
3.  **Create Your Codespace:**
    * Click the green **`< > Code`** button.
    * Select the **"Codespaces"** tab.
    * Click **"Create codespace on main"**.
4.  **Wait for Build:** GitHub will now build your environment. It will have Go installed and be ready for you in 1-2 minutes.

## Your Workflow: How to Complete a Task (All Groups)

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
    * Write your code to solve the problem (using your group's allowed resources).
    * Save the file (`Ctrl+S` or `Cmd+S`).
4.  **Run the test again** (and again) until it passes.
    ```bash
    go test
    ```
    When your solution is correct, you will see a green `PASS` message.

> **Important:** Do not modify the `_test.go` files! Your goal is to make the existing tests pass by editing the *other* `.go` files (like `taskCheck.go` or `task2.go`).

---

## The Tasks (All Groups)

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

## 🏁 Finishing the Experiment (All Groups)

When you have completed all tasks and all tests pass, please "submit" your work by saving it to the repository.

1.  **Push your code:**
    ```bash
    git add .
    git commit -m "Finished all tasks"
    git push origin main
    ```

2.  **Final Step (Check Your Group):**
    * **Non-AI Group:** You are done! Thank you.
    * **⚠️ AI Group Only:** Do not forget to **export your AI chat logs** and **upload them** to the Microsoft Form link at the top of this file.

That's it! Thank you for participating.
