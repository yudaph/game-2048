package main

import (
	"2048/controller"
	"2048/game"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	games := game.NewGameStruct()
	control := controller.NewController(games)
	control.StartGame()

}
