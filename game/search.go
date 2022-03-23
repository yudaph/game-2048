package game

type functionSearch interface {
	search(i, ii int, g *gameStruct) bool
	stop() bool
}

// searchZero to search 0 value, used to spawn new number
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

// compareArena to compare current arena value and compareTo value, used for check lost
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
