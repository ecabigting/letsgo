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
	fmt.Println("\nWelcome to the Number Guessing Game!")
	// fmt.Println("Im thinking of a numbe between 1 and 100")
	// fmt.Println("You have 5 chances to guess the correct number.")
	// fmt.Println("Current Difficulty Level:", gs.DifficultyLevel)
	fmt.Println()
	for {
		difficultyLevelSet := gs.CheckDifficultyLevel()

		if gs.IsPlaying {
			fmt.Println("")
			fmt.Print("Enter your guess:")
		}

		if !difficultyLevelSet {
			fmt.Print("Enter your choice:")
		}
		// Read the next line of input
		if scanner.Scan() {
			input := scanner.Text()

			// Check for exit command
			if strings.TrimSpace(input) == "exit" || strings.TrimSpace(input) == "quit" {
				fmt.Println("Exiting the game thanks for playing!")
				break
			}

			// Check for game status command
			if strings.TrimSpace(input) == "status" {
				gs.ShowGameStatus()
			}

			if gs.IsPlaying {
				guess, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Invalid guess!")
				} else {

					status := gs.EvaluateUserGuess(guess)

					if status {
						gs.SetDifficultyLevel(0)
					} else {
						if gs.AvailableGuesses > 0 {
							fmt.Println("Try again, you still have", gs.AvailableGuesses, "guesses")
						} else {
							fmt.Println("No more Available Guesses!")
							fmt.Println("The number was,", gs.CurrentNumberToGuess)
							fmt.Println(">>>>>> Play again? <<<<<<<")
							gs.IsPlaying = false
							gs.DifficultyLevel = 0
						}
					}
				}
			}

			// check if difficulty level is not yet set,
			// then the latest input is for the difficulty level
			if !difficultyLevelSet {
				level, err := strconv.Atoi(input)
				if err != nil && strings.TrimSpace(input) != "status" {
					fmt.Println("Invalid level selection!")
				} else {
					gs.SetDifficultyLevel(level)
					fmt.Println("Awesome! You selected", gs.GetGameDifficultyAsString(), "level")
					fmt.Println("Lets start the game!")
					fmt.Println("Guess the number between 1 to 100")
					fmt.Println("")
				}
			}

		}

		// Check for errors during input
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			break
		}

	}
}
