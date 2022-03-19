package game

type Action interface {
	search(i, ii int, g *gameStruct) bool
	reset(g *gameStruct) bool
}

type moveLeftStruct struct {
	*searchStruct
	rowIndex int
}

func newMoveLeft() *moveLeftStruct {
	return &moveLeftStruct{rowIndex: 0, searchStruct: newSearch(true)}
}

func (m *moveLeftStruct) reset(games *gameStruct) bool {
	for i := len(m.row); i < games.width; i++ {
		m.row = append(m.row, 0)
	}
	//fmt.Println("RESET", m.rowIndex, m.row)
	games.arena[m.rowIndex] = m.row
	m.rowIndex++
	//fmt.Println(m.rowIndex)
	m.row = []int{}
	m.value = -1
	m.index = -1
	return false
}

type moveRightStruct struct {
	*searchStruct
	rowIndex int
}

func newMoveRight() *moveRightStruct {
	return &moveRightStruct{rowIndex: 0, searchStruct: newSearch(true)}
}

func (m *moveRightStruct) reset(games *gameStruct) bool {
	var row []int
	//fmt.Println(m.rowIndex, m.row)
	for i := len(m.row); i < games.width; i++ {
		row = append(row, 0)
	}
	m.row = append(row, m.row...)
	//fmt.Println("RESET", m.rowIndex, m.row)
	games.arena[m.rowIndex] = m.row
	m.rowIndex++
	//fmt.Println(m.rowIndex)
	m.row = []int{}
	m.value = -1
	m.index = -1
	return false
}

type moveTopStruct struct {
	*searchStruct
	rowIndex int
}

func newMoveTop() *moveTopStruct {
	return &moveTopStruct{rowIndex: 0, searchStruct: newSearch(false)}
}

func (m *moveTopStruct) reset(games *gameStruct) bool {
	//fmt.Println("TOP", m.row)
	length := len(m.row)
	for i := 0; i < games.width; i++ {
		if i < length {
			games.arena[i][m.rowIndex] = m.row[i]
			continue
		}
		games.arena[i][m.rowIndex] = 0
	}
	//fmt.Println("RESET", m.rowIndex, m.row)
	m.rowIndex++
	//fmt.Println(m.rowIndex)
	m.row = []int{}
	m.value = -1
	m.index = -1
	//games.print()
	return false
}

type moveDownStruct struct {
	*searchStruct
	rowIndex int
}

func newMoveDown() *moveDownStruct {
	return &moveDownStruct{rowIndex: 0, searchStruct: newSearch(false)}
}

func (m *moveDownStruct) reset(games *gameStruct) bool {
	length := len(m.row)
	for i := 0; i < games.width; i++ {
		reverse := games.width - 1 - i
		if i < length {
			games.arena[reverse][m.rowIndex] = m.row[length-i-1]
			continue
		}
		games.arena[reverse][m.rowIndex] = 0
	}
	//fmt.Println("RESET", m.rowIndex, m.row)
	m.rowIndex++
	//fmt.Println(m.rowIndex)
	m.row = []int{}
	m.value = -1
	m.index = -1
	return false
}
