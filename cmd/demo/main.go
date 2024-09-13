package main

import (
	"fmt"

	"github.com/NotCoffee418/goscriptoperations"
)

func main() {
	// Register any operations
	goscriptoperations.RegisterOperation(goscriptoperations.Operation{
		Index:       1, // Not literal, just asc sorting order
		Description: "Ask questions",
		Function:    askQuestions,
	})

	// Start listening for user input and execute operations as prompted
	goscriptoperations.StartListening()
}

func askQuestions() {
	q1 := goscriptoperations.PoseBooleanQuestion("Do you like coffee?", true)
	q2 := goscriptoperations.PoseChoiceQuestion("What is your favorite coffee?", []string{"Espresso", "Cappuccino", "Latte", "Americano"}, 1)
	q3 := goscriptoperations.PoseFloatQuestion("How much do you like coffee?", 1.0)
	q4 := goscriptoperations.PoseIntQuestion("How many cups of coffee do you drink a day?", 1)

	fmt.Println(q1, q2, q3, q4)
}
