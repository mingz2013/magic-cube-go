package magic_cube

type Cell int

func (c Cell) Color() int {
	return int(c)
}

func NewCell(color int) Cell {
	return Cell(color)
}

type Side [9]Cell

func (s *Side) Init(color int) {
	for i := 0; i < len(s); i++ {
		s[i] = NewCell(color)
	}
}

func NewSide(color int) Side {
	s := Side{}
	s.Init(color)
	return s
}

type MagicCube [6]Side

func (m *MagicCube) Init() {
	for i := 0; i < len(m); i++ {
		m[i] = NewSide(i)
	}
}

func NewMagicCube() MagicCube {
	m := MagicCube{}
	m.Init()
	return m
}

func (m *MagicCube) Random() {

}

const (
	F = 0 // front
	B = 1 // back
	L = 2 // left
	R = 3 // right
	U = 4 // up
	D = 5 // down
)
