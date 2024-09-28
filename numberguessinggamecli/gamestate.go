package main

type GameState struct {
	DifficultyLevel int
}

func (gs *GameState) InitializeGameState() {
	gs.DifficultyLevel = -1
}
