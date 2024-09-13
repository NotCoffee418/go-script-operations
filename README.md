# Go User Input Utility

Library for settign up CLI with multiple operations in a structured way.

## Features

- Boolean questions (yes/no)
- String input
- Integer input
- Float input
- Multiple choice questions

## Usage

Import the package in your Go project:

```go
import "github.com/NotCoffee418/goscriptoperations"

func main() {
	// Register any operations
	goscriptoperations.RegisterOperation(goscriptoperations.Operation{
		Index:       1, // Not literal, just asc sorting order
		Description: "Example Operation",
		Function:    someFunction,
	})

	// Start listening for user input and execute operations as prompted
	goscriptoperations.StartListening()
}
```

```
# Output
Available Operations:
0. Exit the program
1. Example Operation

Enter operation number: 1

Executing operation 1

Example Operation Completed!
```

## UtilityFunctions

- `PoseBooleanQuestion`: Ask a yes/no question
- `PoseStringQuestion`: Ask for a string input
- `PoseIntQuestion`: Ask for an integer input
- `PoseFloatQuestion`: Ask for a float input
- `PoseChoiceQuestion`: Present a list of choices
