package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameState struct {
	DifficultyLevel      int
	AvailableGuesses     int
	CurrentNumberToGuess int
	UserLastGuess        int
	UserWon              bool
	IsPlaying            bool
}

func (gs *GameState) InitializeGameState() {
	gs.DifficultyLevel = 0
	gs.AvailableGuesses = 0
	gs.UserWon = false
	gs.IsPlaying = false

	// set the number to guess
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	gs.CurrentNumberToGuess = r.Intn(100) + 1
}

func (gs *GameState) CheckDifficultyLevel() bool {
	if gs.DifficultyLevel < 1 || gs.DifficultyLevel > 3 {
		fmt.Println("Please select the difficulty level to continue:")
		fmt.Println("1. Easy (10 Chances)")
		fmt.Println("2. Medium (5 Chances)")
		fmt.Println("3. Hard (3 Chances)")
		fmt.Println("Current level:", gs.DifficultyLevel)
		fmt.Println()
		return false
	}
	return true
}

func (gs *GameState) SetDifficultyLevel(selectedLevel int) bool {
	if selectedLevel >= 1 && selectedLevel <= 3 {
		gs.DifficultyLevel = selectedLevel
		gs.SetGuesses()
		gs.IsPlaying = true
		return true
	}
	return false
}

func (gs *GameState) SetGuesses() {
	switch gs.DifficultyLevel {
	case 1:
		gs.AvailableGuesses = 10
	case 2:
		gs.AvailableGuesses = 5
	case 3:
		gs.AvailableGuesses = 3
	default:
		gs.AvailableGuesses = 0
	}
}

func (gs *GameState) ShowGameStatus() {

	fmt.Println("")
	fmt.Println("-- CURRENT GAME STATUS --")
	fmt.Println("Current Difficulty Level:", gs.DifficultyLevel)
	fmt.Println("Current Available Guesses:", gs.AvailableGuesses)
	fmt.Println("Current Number To Guess:", gs.CurrentNumberToGuess)
	fmt.Println("")
}

func (gs *GameState) EvaluateUserGuess(userGuess int) bool {
	gs.UserLastGuess = userGuess
	gs.AvailableGuesses--

	if userGuess == gs.CurrentNumberToGuess {
		fmt.Println("Good job! Number to guess was:", gs.CurrentNumberToGuess)
		fmt.Println("You win! Remaining guesses:", gs.AvailableGuesses)
		gs.IsPlaying = false
		gs.UserWon = true
		gs.DifficultyLevel = 0
		return true
	}

	if gs.CurrentNumberToGuess > userGuess {
		fmt.Println("Incorrect! Number is greater than", userGuess)
		return false
	}

	if gs.CurrentNumberToGuess < userGuess {
		fmt.Println("Incorrect! Number is less than", userGuess)
		return false
	}
	fmt.Println("Guess the number between 1 to 100")
	return false
}

func (gs *GameState) GetGameDifficultyAsString() string {
	switch gs.DifficultyLevel {
	case 1:
		return "Easy"
	case 2:
		return "Medium"
	case 3:
		return "Hard"
	default:
		return ""
	}
}
