package game

import (
	"testing"
)

// tests several sets of neigbors on a 10x10 game
func TestNeigbors(t *testing.T) {
	b := MakeBoard(10, 10, 10)
	var tests = []struct {
		inputX, inputY int
		outputNeigbors [][2]int
	}{
		// center board
		{4, 3,
			[][2]int{
				{3, 2}, {3, 3}, {3, 4}, {4, 2},
				{4, 4}, {5, 2}, {5, 3}, {5, 4}},
		},
		// top left corner
		{0, 0,
			[][2]int{
				{0, 1}, {1, 0}, {1, 1}},
		},
		// bottom right corner
		{9, 9,
			[][2]int{
				{8, 9}, {9, 8}, {8, 8}},
		},
	}

	for _, test := range tests {
		n := b.neigborCoords(test.inputX, test.inputY)
		expectedMap := make(map[[2]int]bool)

		for _, coord := range test.outputNeigbors {
			expectedMap[coord] = true
		}

		for _, coord := range n {
			if !expectedMap[coord] {
				t.Errorf("(%d, %d) Unexpected Coordinate: %v",
					test.inputX, test.inputY, coord)
			}
			expectedMap[coord] = false
		}

		for coord, b := range expectedMap {
			if b {
				t.Errorf("(%d, %d) Coordinate not created: %v",
					test.inputX, test.inputY, coord)
			}
		}
	}
}
