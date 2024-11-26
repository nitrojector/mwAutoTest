package mwtest

type Cell struct {
	content CellContent
}

type CellContent interface {
	getType() int
}

// CellEmpty is a cell that is empty and can be passed through
type CellEmpty struct{}

func (c CellEmpty) getType() int {
	return 0
}

// CellVoid is a cell that is invalid and should never
// be interacted with
// It is used as a substitute for expanding the board
type CellVoid struct{}

func (c CellVoid) getType() int {
	return 1
}

// CellExit is a cell that represents the exit of the maze
type CellExit struct{}

func (c CellExit) getType() int {
	return 2
}

// CellTeleporter is a cell that represents a teleporter
// Each teleporter can only teleport players to another
// teleporter of the same id
type CellTeleporter struct {
	id int
}

func (c CellTeleporter) getType() int {
	return 3
}
