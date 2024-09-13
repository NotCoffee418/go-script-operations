package goscriptoperations

import "os"

// The exit operation, can be overridden if needed.
var ExitOperation = Operation{
	Index:       -1, // First operation by defalt
	Description: "Exit the program",
	Function:    exitFunction,
}

// Registers internal operations to the operations list
func registerInternalOperations() {
	RegisterOperation(ExitOperation)
}

func exitFunction() {
	os.Exit(0)
}
