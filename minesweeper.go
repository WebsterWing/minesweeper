package main

import (
	"fmt"
	"github.com/WebsterWing/minesweeper/game"
	"github.com/WebsterWing/minesweeper/terminal"
)

func main() {
	fmt.Println("Minesweeper")
	b := game.MakeBoard(10, 10, 10)
	terminal.PrintBoard(b)
	b.FlagClickTile(3, 2)
	b.ClickTile(5, 4)
	terminal.PrintBoard(b)
}
