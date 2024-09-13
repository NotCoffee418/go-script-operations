package goscriptoperations

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// sortOps returns a sorted list of operations
func sortOps() {
	sort.Slice(registeredOperations, func(i, j int) bool {
		return registeredOperations[i].Index < registeredOperations[j].Index
	})
}

// displayOperations displays the available operations
func displayOperations() {
	fmt.Println("Available Operations:")
	sortOps()
	for i, operation := range registeredOperations {
		fmt.Printf("%d. %s\n", i, operation.Description)
	}
}

// requestInput requests input from the user
func requestInput() string {
	fmt.Print("\nEnter operation number: ")
	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	return command
}

// handleCommand handles the command
func handleCommand(command string) {
	if command == "help" {
		displayOperations()
	}

	index, err := strconv.Atoi(strings.TrimSpace(command))
	if err != nil || index < 0 || index >= len(registeredOperations) {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// get the operation
	operation := registeredOperations[index]
	fmt.Printf("\nExecuting operation: %d\n\n", index)
	operation.Function()
}
