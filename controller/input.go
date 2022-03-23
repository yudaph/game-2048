package controller

import "2048/game"

type MoveDirection int

const (
	MoveLeft MoveDirection = iota
	MoveRight
	MoveTop
	MoveDown
)

var (
	dirArr = []game.MoveDirection{game.Left, game.Right, game.Top, game.Down}
)

type Input interface {
	MoveDirection() (MoveDirection, error)
	SetGameSetting() (width, target int, err error)
}

func translateMove(direction MoveDirection) game.Move {
	move := dirArr[direction]
	return game.MoveFactory(move)
}
