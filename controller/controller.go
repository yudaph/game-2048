package controller

import (
	"2048/game"
	"fmt"
	"log"
)

type gameCore interface {
	game.Games
	game.Input
	game.Output
}

type controller struct {
	gameCore
	Input
	Output
}

func NewController(games gameCore, input Input, output Output) *controller {
	return &controller{gameCore: games, Input: input, Output: output}
}

func (c *controller) StartGame() {
	// input game setting : with of arena and target number to win the game
	width, target, err := c.Input.SetGameSetting()
	if err != nil {
		log.Fatal(err)
	}
	// create arena in game
	c.Start(width, target)

	// spawn value in random location in arena
	c.SpawnNewValue(2)

	// print arena, getArena is for request current state of arena
	c.ShowArena(c.GetArena())

	// to looping through the game
	c.gameProgress()
}

func (c *controller) gameProgress() {
	// get desire move direction
	direction, err := c.Input.MoveDirection()
	if err != nil {
		fmt.Println(err)
		c.gameProgress()
		return
	}

	// move to desired direction,
	// translateMove is for translate controller.MoveDirection (get from c.MoveDirection) to game.move
	c.gameCore.MoveTo(translateMove(direction))

	// check is win state achieve
	if c.IsWin() {
		// display win
		c.Output.ShowWin(c.GetArena())
		return
	}

	// spawn new value in random position
	c.SpawnNewValue(2)

	// check is lost achieve
	if c.IsLost() {
		// display lost
		c.ShowLost(c.GetArena())
		return
	}

	// show current state of arena
	c.ShowArena(c.GetArena())

	// call gameProgress again to repeat the iteration
	c.gameProgress()
}
