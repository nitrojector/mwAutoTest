package mwtest

import (
	"fmt"
	"testing"
)

func TestBoard(t *testing.T) {
	c0 := CellTeleporter{id: 0}
	c1 := CellTeleporter{id: 1}
	c0.other = &c1
	c1.other = &c0

	fmt.Println(c0.getType())
	fmt.Println(c1.other.getType())
}
