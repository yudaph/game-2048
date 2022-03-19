package controller

import "2048/game"

type controller struct {
	game.Games
}

func NewController(games game.Games) *controller {
	return &controller{Games: games}
}

func (c *controller) StartGame() {
	c.Start(4, 8)
}

// open close principle
