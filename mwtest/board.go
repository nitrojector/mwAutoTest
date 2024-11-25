package mwtest

import (
	"fmt"
)

type Board struct {
	width, height int
	cells         [][]Cell
	walls         [][]Wall
	// Even indices are vertical (between left-right cells)
	// Odd indices are horizontal (between top-bottom cells)
	// vertical walls have an additional element, but it should not be used
}

func (b *Board) getCell(x, y int) (*Cell, error) {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return nil, fmt.Errorf("Cell (%d, %d) is out of bounds", x, y)
	}
	return &b.cells[x][y], nil

}

type Wall struct {
	exist bool
	// Whether a wall exists or not

	passable bool
	// Walls are not passable iff they bound the board
}
