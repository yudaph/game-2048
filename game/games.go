package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Games interface {
	Start(width, target int)
}

type gameStruct struct {
	width  int
	target int
	arena  [][]int
	win    bool
}

func NewGameStruct() *gameStruct {
	games := new(gameStruct)
	return games
}

func (g *gameStruct) Start(width, target int) {
	g.width = width
	g.arena = make([][]int, width)
	g.win = false
	g.target = target

	for i := range g.arena {
		var col []int
		for ii := 0; ii < width; ii++ {
			col = append(col, 0)
		}
		g.arena[i] = col
	}
	g.spawnValue(2)
	g.cycle()
}

func (g *gameStruct) cycle() {
	g.clean()
	g.print()
	char, err := g.input()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	//fmt.Println(char)
	g.inputProcess(char)
	//fmt.Println(char)
	if !g.win {
		g.spawnValue(2)
		if g.checkLost() {
			g.clean()
			fmt.Println("You Lost")
			g.print()
			return
		}
		g.cycle()
	} else {
		g.printWin()
	}
}

func (g *gameStruct) checkLost() bool {
	newG := &gameStruct{}

	//copy arena to newG
	newG.width = g.width
	newG.arena = make([][]int, g.width)
	for i := range g.arena {
		newG.arena[i] = make([]int, g.width)
		copy(newG.arena[i], g.arena[i])
	}

	//fmt.Println(newG.arena, g.arena)
	newG.actionArena(newMoveLeft())
	compare := newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}
	//fmt.Println(newG.arena, g.arena)

	newG.actionArena(newMoveRight())
	compare = newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}
	//fmt.Println(newG.arena, g.arena)

	newG.actionArena(newMoveTop())
	compare = newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}
	//fmt.Println(newG.arena, g.arena)

	newG.actionArena(newMoveDown())
	compare = newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}
	//fmt.Println(newG.arena, g.arena)

	return true
}

func (g *gameStruct) printWin() {
	g.clean()
	fmt.Println("You win")
	g.print()
}

func (g *gameStruct) input() (char rune, err error) {
	reader := bufio.NewReader(os.Stdin)
	char, _, err = reader.ReadRune()
	if err != nil {
		return
	}
	return
}

func (g *gameStruct) inputProcess(char rune) {
	var move Action
	switch char {
	case 97 | 65:
		move = newMoveLeft()
	case 100 | 68:
		move = newMoveRight()
	case 119 | 87:
		move = newMoveTop()
	case 115 | 83:
		move = newMoveDown()
	}
	g.actionArena(move)
}

func (g *gameStruct) clean() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}

func (g *gameStruct) print() {
	for _, v := range g.arena {
		fmt.Println(v)
	}
}

func (g *gameStruct) searchArena(f functionSearch) {
	for i := 0; i < g.width; i++ {
		for ii := 0; ii < g.width; ii++ {
			found := f.search(i, ii, g)
			if found {
				break
			}
		}
		if f.stop() {
			break
		}
	}
}

func (g *gameStruct) actionArena(f Action) {
	for i := 0; i < g.width; i++ {
		for ii := 0; ii < g.width; ii++ {
			f.search(i, ii, g)
		}
		f.reset(g)
	}
}

func (g *gameStruct) spawnValue(numberOfValue int) {
	s := newSearchZero()
	g.searchArena(s)
	if !s.found {
		return
	}
	rand.Seed(time.Now().UnixNano())
	i, ii := -1, -1
	for {
		i, ii = rand.Intn(g.width), rand.Intn(g.width)
		if g.arena[i][ii] == 0 {
			g.arena[i][ii] = 2
			break
		}
	}
	numberOfValue--
	if numberOfValue > 0 {
		g.spawnValue(numberOfValue)
	}
}
