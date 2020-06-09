package terminal

import (
	"fmt"
	"github.com/WebsterWing/minesweeper/game"
)

const (
	// grid pieces
	horiz     = "\u2500"
	horizUp   = "\u2534"
	horizDown = "\u252c"
	vert      = "\u2502"
	vertRight = "\u251c"
	vertLeft  = "\u2524"
	downRight = "\u250c"
	downLeft  = "\u2510"
	upRight   = "\u2514"
	upLeft    = "\u2518"
	cross     = "\u253c"

	// individual tiles
	coveredTile = "\u2588\u2588"
	flagTile    = "\U0001f6a9"
	explosion   = "\u0489"
)

func PrintBoard(board *game.Board) {
	h, w := board.GetDimentions()

	fmt.Print("   ")
	for i := 0; i < w; i++ {
		fmt.Printf("%3d", i)
	}

	fmt.Println()

	for i := 0; i < h; i++ {
		fmt.Print("   ")
		if i == 0 {
			printGridTop(w)
		} else {
			printGridDivider(w)
		}
		fmt.Printf("\n%2d ", i)

		for j := 0; j < w; j++ {
			fmt.Print(vert)
			state, num := board.GetTile(j, i)
			switch state {
			case game.Covered:
				fmt.Print(coveredTile)
			case game.Flagged:
				fmt.Print(flagTile)
			case game.Mined:
				fmt.Print(explosion)
			case game.Uncovered:
				if num > 0 {
					fmt.Printf("%2d", num)
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println(vert)
	}

	fmt.Print("   ")
	printGridBottom(w)
	fmt.Println()
}

func printGridTop(cells int) {
	fmt.Print(downRight, horiz, horiz)

	for i := 1; i < cells-1; i++ {
		fmt.Print(horizDown, horiz, horiz)
	}

	fmt.Print(horizDown, horiz, horiz, downLeft)
}

func printGridDivider(cells int) {
	fmt.Print(vertRight, horiz, horiz)

	for i := 1; i < cells-1; i++ {
		fmt.Print(cross, horiz, horiz)
	}

	fmt.Print(cross, horiz, horiz, vertLeft)
}

func printGridBottom(cells int) {
	fmt.Print(upRight, horiz, horiz)

	for i := 1; i < cells-1; i++ {
		fmt.Print(horizUp, horiz, horiz)
	}

	fmt.Print(horizUp, horiz, horiz, upLeft)
}
