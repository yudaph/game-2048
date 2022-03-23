package controller

type Output interface {
	ShowArena([][]int)
	ShowWin([][]int)
	ShowLost([][]int)
}
