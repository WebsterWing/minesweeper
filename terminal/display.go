package terminal

import (
	"fmt"
	"github.com/WebsterWing/minesweeper/game"
)

const (
	horiz     = '\u2500'
	horizUp   = '\u2534'
	horizDown = '\u252c'
	vert      = '\u2502'
	vertRight = '\u251c'
	vertLeft  = '\u2524'
	downRight = '\u250c'
	downLeft  = '\u2510'
	upRight   = '\u2514'
	upLeft    = '\u2518'
	cross     = '\u253c'
)

var board *game.Board

func Setup(h, w, mines int) {
	board = game.MakeBoard(h, w, mines)
}

func printBoard() {
	h, w := board.GetDimentions()

	fmt.Print("   ")
	for i := 0; i < w; i++ {
		fmt.Printf("%3d", i)
	}

	for i := 0; i < h; i++ {
		fmt.Print("   ")
		fmt.Printf("\n %v ", rune(byte(i)-'A'))
		printGridDivider(w)
	}

}

func printGridTop(cells int) {
	fmt.Print(downRight, horiz)

	for i := 1; i < cells-1; i++ {
		fmt.Print(horizDown, horiz)
	}

	fmt.Print(horizDown, horiz, downLeft)
}

func printGridDivider(cells int) {
	fmt.Print(vertRight, horiz)

	for i := 1; i < cells-1; i++ {
		fmt.Print(cross, horiz)
	}

	fmt.Print(cross, horiz, vertLeft)
}

func printGridBottom(cells int) {
	fmt.Print(upRight, horiz)

	for i := 1; i < cells-1; i++ {
		fmt.Print(horizUp, horiz)
	}

	fmt.Print(horizUp, horiz, upLeft)
}
