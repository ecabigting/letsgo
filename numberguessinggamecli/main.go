package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	gs := GameState{}
	gs.InitializeGameState()
	// Starting up
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Im thinking of a number between 1 and 100")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println("Currend Difficulty Level:", gs.DifficultyLevel)

	for {
		if gs.DifficultyLevel < 1 || gs.DifficultyLevel > 3 {
			fmt.Println("Please select the difficulty level:")
			fmt.Println("1. Easy (10 Chances)")
			fmt.Println("2. Medium (5 Chances)")
			fmt.Println("3. Hard (3 Chances)")
			fmt.Println("Currend :", gs.DifficultyLevel)
		}

		// Print a prompt
		fmt.Print("> ")
		// Read the next line of input
		if scanner.Scan() {
			input := scanner.Text()

			// Check for exit command
			if strings.TrimSpace(input) == "exit" {
				fmt.Println("Exiting the game thanks for playing!")
				break
			}

			// check if user set a difficultyLevel
			if gs.DifficultyLevel < 1 || difficultyLevel > 3 {
				selectedLevel := strconv.Itoa(strings.TrimSpace(input))
				fmt.Println(selectedLevel)
			}

			// Print the result
			fmt.Println(input)
		}

		// Check for errors during input
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			break
		}
	}
}
