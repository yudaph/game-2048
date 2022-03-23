package output

import (
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}

func printArena(arena [][]int) {
	for _, v := range arena {
		fmt.Println(v)
	}
}
