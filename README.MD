# Grade Calculator

This project is written in Go and is used to calculate a student's final grade based on their assignments, exams, and essays.

# Requirements

- Go 1.24.4 or newer
- Git to clone the repository

# Files

- `README.MD` - Info about the project and how to run it
- `go.mod` - Contains the Go module information
- `grade_calculator.go` - Main code that calculates grades
- `grade_calculator_test.go` - Tests for checking if the calculator works correctly
# How It Works

The calculator takes in grades from three categories: assignments, exams, and essays.

It calculates the average for each category and then uses their weights to calculate the final grade. The final number is converted into a letter grade

# Running the Project

First, clone the repository and open it in VS Code or another editor.

Open a terminal and go into the grade-calculator folder:

`cd grade-calculator`

To run the tests:

`go test ./...`

To see more details about the tests:

`go test -v ./...`

# Code Coverage

You can also check how much of the code is being tested.

Run:

`go test -coverprofile=coverage.out ./...`

Then:

`go tool cover -func=coverage.out`

To see the coverage in a browser:

`go tool cover -html=coverage.out`

## Notes

The project uses Go's built-in testing package. If changes are made to the calculator, the tests can be run again to check if everything still works.