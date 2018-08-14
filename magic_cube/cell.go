package magic_cube

type Cell int

func (c Cell) Color() int {
	return int(c) / 10
}

func NewCell(i int) Cell {
	return Cell(i)
}
