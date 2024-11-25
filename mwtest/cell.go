package mwtest

type Cell struct {
	content CellContent
}

type CellContent interface {
	getType() int
}

type CellEmpty struct{}

func (c CellEmpty) getType() int {
	return 0
}

type CellVoid struct{}

func (c CellVoid) getType() int {
	return 1
}

type CellTeleporter struct {
	id int
}

func (c CellTeleporter) getType() int {
	return 2
}

type CellDarkRegion struct{}

func (c CellDarkRegion) getType() int {
	return 3
}
