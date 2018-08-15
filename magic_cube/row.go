package magic_cube

type Row [3]Cell

func (r *Row) Reverse() *Row {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func NewRow(c1, c2, c3 Cell) *Row {
	return &Row{c1, c2, c3}
}

func (r *Row) Items() (Cell, Cell, Cell) {
	return r[0], r[1], r[2]
}
