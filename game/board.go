package game

import (
	"math/rand"
	"time"
)

const (
	Covered TileState = iota
	Uncovered
	Flagged
	Mined
)

type TileState int

type tile struct {
	state     TileState
	mineCount int
	mine      bool
}

type Board struct {
	tiles         []tile
	height, width int
	mines, flags  int
	started       bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MakeBoard(h, w, mineCount int) *Board {
	b := new(Board)
	b.tiles = make([]tile, h*w)
	b.height, b.width = h, w
	b.mines = mineCount

	return b
}

func (b *Board) indexOf(x, y int) int {
	if x < 0 || y < 0 || x >= b.width || y >= b.height {
		return -1
	}

	return y*b.width + x
}

func (b *Board) ValidTile(x, y int) bool {
	return b.indexOf(x, y) >= 0
}

func (b *Board) ClickTile(x, y int) bool {
	i := b.indexOf(x, y)

	if !b.started {
		b.setupFirstClick(x, y)
		b.started = true
	}

	if i < 0 {
		return false
	}

	if b.tiles[i].mine {
		b.tiles[i].state = Mined
		return false
	}

	b.revealTile(x, y)
	return true
}

// Sets up the board after the first click so the clicked tile
// doesn't have a mine, and neither do any of the neigbors
func (b *Board) setupFirstClick(x, y int) {
	h, w := b.GetDimentions()
	neigbors := b.neigborCoords(x, y)
	numClearTiles := len(neigbors) + 1

	// Always leave numClearTiles clear tiles at the end
	for i := 0; i < b.mines && i < h*w-numClearTiles; i++ {
		b.tiles[i].mine = true
	}

	swap := func(i, j int) {
		b.tiles[i], b.tiles[j] = b.tiles[j], b.tiles[i]
	}

	rand.Shuffle(h*w-numClearTiles, swap)

	// swap clear area with space at end of track
	i := h*w - numClearTiles // first clear tile
	swap(b.indexOf(x, y), i)
	i++

	for j := 0; i < h*w; i, j = i+1, j+1 {
		swap(b.indexOf(neigbors[j][0], neigbors[j][1]), i)
	}
}

func (b *Board) neigborCoords(x, y int) [][2]int {
	ret := make([][2]int, 0, 8)

	for i := 0; i < 9; i++ {
		xAdj := (i % 3) - 1
		yAdj := i/3 - 1

		if x == 0 && y == 0 {
			continue
		}

		if b.indexOf(x+xAdj, y+yAdj) >= 0 {
			ret = append(ret, [2]int{x + xAdj, y + yAdj})
		}
	}

	return ret
}

func (b *Board) revealTile(x, y int) {
	tile := &b.tiles[b.indexOf(x, y)]
	if tile.state != Covered {
		return
	}

	nearbyMines := 0
	neigbors := b.neigborCoords(x, y)

	// count neigboring mines
	for _, coord := range neigbors {
		if b.tiles[b.indexOf(coord[0], coord[1])].mine {
			nearbyMines++
		}
	}

	tile.mineCount = nearbyMines
	tile.state = Uncovered

	// reveal surrounding tiles if no neigbors are mines
	if nearbyMines == 0 {
		for _, coord := range neigbors {
			b.revealTile(coord[0], coord[1])
		}
	}
}

func (b *Board) FlagClickTile(x, y int) {
	i := b.indexOf(x, y)
	if i < 0 || !b.started {
		return
	}

	switch b.tiles[i].state {
	case Covered:
		b.tiles[i].state = Flagged
	case Flagged:
		b.tiles[i].state = Covered
	}
}

func (b *Board) GetDimentions() (height, width int) {
	return b.height, b.width
}

func (b *Board) GetTile(x, y int) (state TileState, surrounding int) {
	i := b.indexOf(x, y)
	tile := &b.tiles[i]
	return tile.state, tile.mineCount
}
