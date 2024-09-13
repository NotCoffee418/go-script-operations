package goscriptoperations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PoseBooleanQuestion poses a boolean question to the user
// returns the boolean value
func PoseBooleanQuestion(question string, defaultResponse bool) bool {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s [%s/%s]: ", question, map[bool]string{true: "Y", false: "N"}[defaultResponse], map[bool]string{true: "n", false: "y"}[defaultResponse])
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "" {
			return defaultResponse
		}

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}

		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}

// PoseStringQuestion poses a string question to the user
func PoseStringQuestion(question string, defaultResponse string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s]: ", question, defaultResponse)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	if response == "" {
		return defaultResponse
	}

	return response
}

// PoseIntQuestion poses an int question to the user
// returns the int value
func PoseIntQuestion(question string, defaultResponse int) int {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s [%d]: ", question, defaultResponse)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response == "" {
			return defaultResponse
		}

		parsedResponse, err := strconv.Atoi(response)
		if err == nil {
			return parsedResponse
		}

		fmt.Println("Invalid input. Please enter a valid integer.")
	}
}

// PoseFloatQuestion poses a float question to the user
// returns the float value
func PoseFloatQuestion(question string, defaultResponse float64) float64 {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s [%f]: ", question, defaultResponse)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response == "" {
			return defaultResponse
		}

		parsedResponse, err := strconv.ParseFloat(response, 64)
		if err == nil {
			return parsedResponse
		}

		fmt.Println("Invalid input. Please enter a valid float.")
	}
}

// PoseChoiceQuestion poses a choice question to the user
// returns the index of the choice
func PoseChoiceQuestion(question string, choices []string, defaultChoice int) int {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s (default: %d)\n", question, defaultChoice+1)
		for i, choice := range choices {
			fmt.Printf("%d. %s\n", i+1, choice)
		}
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response == "" {
			return defaultChoice
		}

		parsedResponse, err := strconv.Atoi(response)
		if err == nil && parsedResponse >= 1 && parsedResponse <= len(choices) {
			return parsedResponse - 1
		}

		fmt.Printf("Invalid input. Please enter a number between 1 and %d.\n", len(choices))
	}
}
