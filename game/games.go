package game

import (
	"math/rand"
	"time"
)

// Games is core interface
type Games interface {
	Start(width, target int)
	SpawnNewValue(numberOfValue int)
}

// Input is interface to get input to the game
type Input interface {
	MoveTo(action Move)
}

// Output is interface display the game
type Output interface {
	GetArena() [][]int
	IsWin() bool
	IsLost() bool
}

// gameStruct is core struct, where the game is actually saved and processed
type gameStruct struct {
	width  int
	target int
	arena  [][]int
	win    bool
}

// NewGameStruct create empty gameStruct
func NewGameStruct() *gameStruct {
	return &gameStruct{}
}

// Start to create game arena and set win value to achieve
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
}

// SpawnNewValue to add number of values to arena in random position
func (g *gameStruct) SpawnNewValue(numberOfValue int) {
	// search is there any 0 (zero) value in arena
	s := newSearchZero()
	g.searchArena(s)

	// return if there is no 0 (zero) value
	if !s.found {
		return
	}

	// init random index to spawn value in random position
	rand.Seed(time.Now().UnixNano())
	i, ii := -1, -1
	// loop random index
	for {
		//if random index generated value equals to 0 (zero) spawn new number there and broke loop
		i, ii = rand.Intn(g.width), rand.Intn(g.width)
		if g.arena[i][ii] == 0 {
			g.arena[i][ii] = 2
			break
		}
	}

	// if we need to spawn another value, call this function again
	numberOfValue--
	if numberOfValue > 0 {
		g.SpawnNewValue(numberOfValue)
	}
}

// GetArena to get current state of arena
func (g *gameStruct) GetArena() [][]int {
	return g.arena
}

// IsWin to get current state of win
func (g *gameStruct) IsWin() bool {
	return g.win
}

// IsLost to check if there is no possible move
func (g *gameStruct) IsLost() bool {
	newG := &gameStruct{}

	//copy arena to newG
	newG.width = g.width
	newG.arena = make([][]int, g.width)
	for i := range g.arena {
		newG.arena[i] = make([]int, g.width)
		copy(newG.arena[i], g.arena[i])
	}

	//if newG move left change arena return false
	newG.actionArena(newMoveLeft())
	compare := newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}

	//if newG move right change arena return false
	newG.actionArena(newMoveRight())
	compare = newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}

	//if newG move top change arena return false
	newG.actionArena(newMoveTop())
	compare = newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}

	//if newG move down change arena return false
	newG.actionArena(newMoveDown())
	compare = newCompareArena(newG.arena)
	g.searchArena(compare)
	if compare.notEquals {
		return false
	}

	// if all move didn't change arena return true
	return true
}

// MoveTo to move value inside the arena to desire direction
func (g *gameStruct) MoveTo(action Move) {
	g.actionArena(action)
}

// private function

// function to search or compare value in arena column
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

// function to move value in arena
func (g *gameStruct) actionArena(f Move) {
	for i := 0; i < g.width; i++ {
		for ii := 0; ii < g.width; ii++ {
			f.checkCell(i, ii, g)
		}
		f.reposition(g)
	}
	f.reset()
}
