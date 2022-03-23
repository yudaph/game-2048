package input

import (
	"2048/controller"
	"errors"
	"fmt"
	"strconv"
)

type inputStruct struct {
}

func NewInputStruct() *inputStruct {
	return &inputStruct{}
}

func (i *inputStruct) SetGameSetting() (width, target int, err error) {
	fmt.Println("input width arena")
	char, err := input(false)
	if err != nil {
		err = errors.New("input width arena not valid")
		return
	}
	width, err = strconv.Atoi(char)
	if err != nil {
		err = errors.New("input width must only contain number")
		return
	}

	fmt.Println("input target number")
	char, err = input(false)
	if err != nil {
		err = errors.New("input target not valid")
		return
	}
	target, err = strconv.Atoi(char)
	if err != nil {
		err = errors.New("input target must only contain number")
		return
	}

	return
}

func (i *inputStruct) MoveDirection() (move controller.MoveDirection, err error) {
	str, err := input(true)
	if err != nil {
		err = errors.New("input move not valid")
		return
	}

	return getDirectionFromStr(str)
}

func getDirectionFromStr(str string) (move controller.MoveDirection, err error) {
	switch str {
	case "a", "A":
		move = controller.MoveLeft
	case "d", "D":
		move = controller.MoveRight
	case "w", "W":
		move = controller.MoveTop
	case "s", "S":
		move = controller.MoveDown
	default:
		err = errors.New("input move not valid, please input a,s,d, or w")
	}
	return
}
