package main

import (
	"2048/controller"
	"2048/game"
	"2048/input"
	"2048/output"
)

func main() {
	games := game.NewGameStruct()

	// change in for new input process
	in := input.NewInputStruct()

	// change out for new output process
	out := output.NewOutputStruct()
	control := controller.NewController(games, in, out)
	control.StartGame()
}
