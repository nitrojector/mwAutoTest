package mwtest

import (
	"fmt"
)

// Represent the board and only the board, not elements on the board
type Board struct {
	width, height int
	cells         [][]Cell
	walls         [][]int
	// (-1) Wall that does not allow passthrough
	// (0) No wall
	// (1) Wall that allows passthrough
	// Even indices are vertical (between left-right cells)
	// Odd indices are horizontal (between top-bottom cells)
	// 0 1 2 3 4 = x
	// * | * | * // y = 0
	// -   -   - // y = 0
	// * | * | * // y = 1
	// -   -   - // y = 1 // These walls should be set to -1
}

func NewBoard(x, y int) *Board {
	b := Board{width: x, height: y}
	b.cells = make([][]Cell, x)
	b.walls = make([][]int, 2*x-1)
	for i := 0; i < x; i++ {
		b.cells[i] = make([]Cell, y)
	}
	for i := 0; i < 2*x-1; i++ {
		b.walls[i] = make([]int, y)
	}
	return &b
}

func (b *Board) getCell(x, y int) (*Cell, error) {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return nil, fmt.Errorf("Cell (%d, %d) is out of bounds", x, y)
	}
	return &b.cells[x][y], nil

}

// Returns walls surrounding the cell at (x, y)
// The order of the walls is as follows:
// 0: Top
// 1: Right
// 2: Bottom
// 3: Left
func (b *Board) getWalls(x, y int) ([]int, error) {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return nil, fmt.Errorf("Cell (%d, %d) is out of bounds", x, y)
	}

	top, right, bottom, left := -1, -1, -1, -1

	if y != 0 {
		top = b.walls[2*x][y-1]
	}
	if x != b.width-1 {
		right = b.walls[2*x+1][y]
	}
	if y != b.height-1 {
		bottom = b.walls[2*x][y]
	}
	if x != 0 {
		left = b.walls[2*x-1][y]
	}

	return []int{top, right, bottom, left}, nil
}
