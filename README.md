# 2048 Game
create a game look a like 2048, 
but with flexibility to customize the target value for win and the with of the arena

## how to use
Clone from this repository
```bash
git clone https://github.com/yudaph/game-2048.git
```

Run the game with existing input and output which is command prompt
```bash
go run main.go
```
how to control the game :
- `a` to move `left`
- `w` to move `top`
- `s` to move `down`
- `d` to move `right`

## development
this game is open for development such as change the input and/or output method

### modify input
there is two ways to modify input, we can directly connect to `game` `Input` interface
```go
type Input interface {
	MoveTo(action Move)
}
```
or using `controller` `Input` interface (`recommended`)
```go
type Input interface {
	MoveDirection() (MoveDirection, error)
	SetGameSetting() (width, target int, err error)
}
```
### modify output
and also there is two ways for modifying output, directly connect via `game` `output` interface
```go
type Output interface {
	GetArena() [][]int
	IsWin() bool
	IsLost() bool
}
```
or using `controller` `Output` interface (`recommended`)
```go
type Output interface {
	ShowArena([][]int)
	ShowWin([][]int)
	ShowLost([][]int)
}
```

don't forget to modify `main.go` if we decided to change via `controller` interface
```go
func main() {
	games := game.NewGameStruct()
	
	// change in for new input process
	in := input.NewInputStruct()
	
	// change out for new output process
	out := output.NewOutputStruct()
	control := controller.NewController(games, in, out)
	control.StartGame()
}
```

### modify game flow
we can also modify game flow via `controller` `gameProgress` function
```go
func (c *controller) gameProgress() {
	// get desire move direction
	direction, err := c.Input.MoveDirection()
	if err != nil {
		fmt.Println(err)
		c.gameProgress()
		return
	}

	// move to desired direction,
	// translateMove is for translate controller.MoveDirection (get from c.MoveDirection) to game.move
	c.gameCore.MoveTo(translateMove(direction))

	// check is win state achieve
	if c.IsWin() {
		// display win
		c.Output.ShowWin(c.GetArena())
		return
	}

	// spawn new value in random position
	c.SpawnNewValue(2)

	// check is lost achieve
	if c.IsLost() {
		// display lost
		c.ShowLost(c.GetArena())
		return
	}

	// show current state of arena
	c.ShowArena(c.GetArena())

	// call gameProgress again to repeat the iteration
	c.gameProgress()
}
```

# future development
- `-`