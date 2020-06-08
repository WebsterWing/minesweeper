package game

import (
	"math/rand"
)

const (
	Covered = iota
	Uncovered
	Flagged
)

type tile struct {
	state     int
	mineCount int
	mine      bool
}

type Board struct {
	tiles         []tile
	height, width int
	mines, flags  int
}

func MakeBoard(h, w, mineCount int) *Board {
	b := new(Board)
	b.tiles = make([]tile, h*w)
	b.height, b.width = h, w
	b.mines = mineCount

	for i := 0; i < mineCount && i < h*w-9; i++ {
		b.tiles[i].mine = true
	}

	rand.Shuffle(h*w, func(i, j int) {
		b.tiles[i], b.tiles[j] = b.tiles[j], b.tiles[i]
	})

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
	if i < 0 || b.tiles[i].mine {
		return false
	}

	b.revealTile(x, y)
	return true
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
	if tile.state == Flagged {
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
	if i < 0 {
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

func (b *Board) GetTile(x, y int) (state, surrounding int) {
	i := b.indexOf(x, y)
	tile := &b.tiles[i]
	return tile.state, tile.mineCount
}
