package magic_cube

// 第一种思路，按每个面分，一个面9个cell，一共6个面

type Cell int

func (c Cell) Color() int {
	return int(c) / 10
}

func NewCell(i int) Cell {
	return Cell(i)
}

type Side [9]Cell

func (s *Side) Init(color int) {
	for i := 0; i < len(s); i++ {
		s[i] = NewCell(color*10 + i)
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

// 前F后B左L右R上U下D
const (
	F = 0 // front
	B = 1 // back
	L = 2 // left
	R = 3 // right
	U = 4 // up
	D = 5 // down
)

func NewMagicCube() MagicCube {
	m := MagicCube{}
	m.Init()
	return m
}

func (m *MagicCube) Random() {

}

// 读取不同面

func (m *MagicCube) FSide() Side {
	return m[F]
}

func (m *MagicCube) BSide() Side {
	return m[B]
}

func (m *MagicCube) LSide() Side {
	return m[L]
}

func (m *MagicCube) RSide() Side {
	return m[R]
}

func (m *MagicCube) USide() Side {
	return m[U]
}

func (m *MagicCube) DSide() Side {
	return m[D]
}

// 一系列操作，旋转操作
func (m *MagicCube) U() {
	// 上层做顺时针90度旋转
	// 涉及前后左右四面的上层，涉及U的整个变换

}

func (m *MagicCube) UR() {
	// 上层做逆时针90度旋转
}

func (m *MagicCube) U2() {
	// 上层顺时针旋转180度，或，上层逆时针旋转180度
}

func (m *MagicCube) R() {
	// 右侧
}

func (m *MagicCube) RR() {

}

func (m *MagicCube) R2() {

}

func (m *MagicCube) F() {
	// 前面
}
func (m *MagicCube) FR() {

}

func (m *MagicCube) F2() {

}

func (m *MagicCube) L() {

}

func (m *MagicCube) LR() {

}

func (m *MagicCube) L2() {

}

func (m *MagicCube) B() {

}

func (m *MagicCube) BR() {

}

func (m *MagicCube) B2() {

}

func (m *MagicCube) D() {

}

func (m *MagicCube) DR() {

}

func (m *MagicCube) D2() {

}
