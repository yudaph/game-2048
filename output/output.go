package output

import "fmt"

type outputStruct struct {
}

func NewOutputStruct() *outputStruct {
	return &outputStruct{}
}

func (o *outputStruct) ShowArena(arena [][]int) {
	clear()
	printArena(arena)
}

func (o *outputStruct) ShowWin(arena [][]int) {
	clear()
	fmt.Println("You Win")
	printArena(arena)
}

func (o *outputStruct) ShowLost(arena [][]int) {
	clear()
	fmt.Println("You Lost")
	printArena(arena)
}
