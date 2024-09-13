package goscriptoperations

var registeredOperations []Operation = []Operation{}

// RegisterOperation registers an operation to the operations list
func RegisterOperation(operation Operation) {
	registeredOperations = append(registeredOperations, operation)
}

// RegisterOperations registers a list of operations to the operations list
func RegisterOperations(operations []Operation) {
	registeredOperations = append(registeredOperations, operations...)
}

// Start displays the CLI and listens for user input
func StartListening() {
	registerInternalOperations()
	for {
		displayOperations()
		command := requestInput()
		handleCommand(command)
	}
}
