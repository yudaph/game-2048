package game

import "log"

type MoveDirection int

const (
	Left MoveDirection = iota
	Right
	Top
	Down
)

type Move interface {
	checkCell(i, ii int, g *gameStruct)
	reposition(g *gameStruct)
	reset()
}

type move struct {
	isHorizontalSearch bool
	row                []int
	index              int
	value              int
	rowIndex           int
}

func newMove(isHorizontalSearch bool) *move {
	return &move{index: -1, value: -1, isHorizontalSearch: isHorizontalSearch, rowIndex: 0}
}

// MoveFactory to create Move based on direction
func MoveFactory(direction MoveDirection) Move {
	switch direction {
	case Left:
		return newMoveLeft()
	case Right:
		return newMoveRight()
	case Top:
		return newMoveTop()
	case Down:
		return newMoveDown()
	default:
		log.Fatal("error when create move")
		return nil
	}
}

func (m *move) reset() {
	m.row = []int{}
	m.index = -1
	m.value = -1
	m.rowIndex = 0
}

// function to checkCell value and save to another slice if value more than 0
func (m *move) checkCell(i, ii int, games *gameStruct) {

	// get value
	value := games.arena[ii][i]
	if m.isHorizontalSearch {
		value = games.arena[i][ii]
	}

	// compare the current value and value before
	// if value match current value in saved slice (m.row) will be doubled
	// if the doubled value equals or more than games target, update game status win to true
	// then clear value before to -1, so it can be used to search next value
	// end the function
	if m.value == value {
		m.row[m.index] *= 2
		if m.row[m.index] >= games.target {
			games.win = true
		}
		m.value = -1
		return
	}

	// if current value more than 0
	// add current value to m.row slice to be matched with next non-zero value
	// add pointer in m.index to indicate new last value in slice
	if value > 0 {
		m.row = append(m.row, value)
		m.value = value
		m.index++
	}
}

// nextReposition to reset and prepared variable for next row / column iteration
func (m *move) nextReposition() {
	// add rowIndex to move to next row
	m.rowIndex++

	// reset move value
	m.row = []int{}
	m.value = -1
	m.index = -1
}

// moveLeftStruct
type moveLeftStruct struct {
	*move
}

func newMoveLeft() *moveLeftStruct {
	return &moveLeftStruct{move: newMove(true)}
}

func (m *moveLeftStruct) reposition(games *gameStruct) {
	// add zero value in saved slice to match arena width
	for i := len(m.row); i < games.width; i++ {
		m.row = append(m.row, 0)
	}

	// replace current row with saved row, it will make all non-zero value move to left
	games.arena[m.rowIndex] = m.row

	// reset to move to next row
	m.nextReposition()
}

// moveRightStruct
type moveRightStruct struct {
	*move
}

func newMoveRight() *moveRightStruct {
	return &moveRightStruct{move: newMove(true)}
}

func (m *moveRightStruct) reposition(games *gameStruct) {
	//create new variable to match arena width
	var row []int
	for i := len(m.row); i < games.width; i++ {
		row = append(row, 0)
	}
	// add new created variable before non-zero value, it will make all non-zero value move to right side
	m.row = append(row, m.row...)

	// replace current row arena with m.row
	games.arena[m.rowIndex] = m.row

	// reset to move to next row
	m.nextReposition()
}

// moveTopStruct start
type moveTopStruct struct {
	*move
}

func newMoveTop() *moveTopStruct {
	return &moveTopStruct{move: newMove(false)}
}

func (m *moveTopStruct) reposition(games *gameStruct) {
	// place the non-zero value slice (m.row) to each column one by one from the top
	length := len(m.row)
	for i := 0; i < games.width; i++ {
		if i < length {
			games.arena[i][m.rowIndex] = m.row[i]
			continue
		}
		// after all non-zero value slice (m.row) placed, replace other value with 0
		// it will make all non-zero value move to top
		games.arena[i][m.rowIndex] = 0
	}

	// reset to move to next column
	m.nextReposition()
}

// moveDownStruct
type moveDownStruct struct {
	*move
}

func newMoveDown() *moveDownStruct {
	return &moveDownStruct{move: newMove(false)}
}

func (m *moveDownStruct) reposition(games *gameStruct) {
	length := len(m.row)
	for i := 0; i < games.width; i++ {
		reverse := games.width - 1 - i
		if i < length {
			games.arena[reverse][m.rowIndex] = m.row[length-i-1]
			continue
		}
		games.arena[reverse][m.rowIndex] = 0
	}

	// reset to move to next column
	m.nextReposition()
}
