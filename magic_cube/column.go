package magic_cube

type Column [3]Cell

func (c *Column) Reverse() Column {
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	return *c
}
