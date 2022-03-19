package game

type functionSearch interface {
	search(i, ii int, g *gameStruct) bool
	stop() bool
}

type searchStruct struct {
	isHorizontalSearch bool
	row                []int
	index              int
	value              int
}

func newSearch(isHorizontalSearch bool) *searchStruct {
	return &searchStruct{index: -1, value: -1, isHorizontalSearch: isHorizontalSearch}
}

func (b *searchStruct) search(i, ii int, games *gameStruct) bool {
	stopSearch := false
	value := games.arena[ii][i]
	if b.isHorizontalSearch {
		value = games.arena[i][ii]
	}
	//fmt.Println(ii, i, value)
	if b.value == value {
		b.row[b.index] *= 2
		if b.row[b.index] >= games.target {
			games.win = true
		}
		b.value = -1
		return stopSearch
	}
	if value > 0 {
		b.row = append(b.row, value)
		b.value = value
		b.index++
	}
	//fmt.Println(b.row)
	return stopSearch
}

type searchZero struct {
	found bool
}

func newSearchZero() *searchZero {
	return &searchZero{found: false}
}

func (s *searchZero) search(i, ii int, games *gameStruct) bool {
	if games.arena[i][ii] == 0 {
		s.found = true
		return true
	}
	return false
}

func (s *searchZero) stop() bool {
	if s.found == true {
		return true
	}
	return false
}

type compareArena struct {
	notEquals bool
	compareTo [][]int
}

func newCompareArena(compareTo [][]int) *compareArena {
	return &compareArena{notEquals: false, compareTo: compareTo}
}

func (c *compareArena) search(i, ii int, games *gameStruct) bool {
	if games.arena[i][ii] != c.compareTo[i][ii] {
		c.notEquals = true
		return true
	}
	return false
}

func (c *compareArena) stop() bool {
	if c.notEquals == true {
		return true
	}
	return false
}
