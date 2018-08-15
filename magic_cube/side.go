package magic_cube

type Side [9]Cell // 二维的，应当有二维的操作

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

func (s *Side) isBingo() (b bool) {
	color := -1
	for i := 0; i < len(s); i++ {
		c := s[i]
		if color == -1 {
			color = c.Color()
		} else {
			if color != c.Color() {
				return false
			}
		}
	}
	return true
}

func (s *Side) clockwise() {
	// 顺时针90度

	s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8] = s[6], s[3], s[0], s[7], s[4], s[1], s[8], s[5], s[2]

	//log.Println(s)
}

func (s *Side) anticlockwise() {
	// 逆时针90度
	s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8] = s[2], s[5], s[8], s[1], s[4], s[7], s[0], s[3], s[6]
}

func (s *Side) X0() (c1, c2, c3 Cell) {
	return s[0], s[3], s[6]
}

func (s *Side) setX0(c1, c2, c3 Cell) {
	s[0], s[3], s[6] = c1, c2, c3
}

func (s *Side) setX1(c1, c2, c3 Cell) {
	s[1], s[4], s[7] = c1, c2, c3
}

func (s *Side) X1() (c1, c2, c3 Cell) {
	return s[1], s[4], s[7]
}

func (s *Side) setX2(c1, c2, c3 Cell) {
	s[2], s[5], s[8] = c1, c2, c3
}

func (s *Side) X2() (c1, c2, c3 Cell) {
	return s[2], s[5], s[8]
}

func (s *Side) setY0(c1, c2, c3 Cell) {
	s[0], s[1], s[2] = c1, c2, c3
}

func (s *Side) Y0() (c1, c2, c3 Cell) {
	return s[0], s[1], s[2]
}

func (s *Side) setY1(c1, c2, c3 Cell) {
	s[3], s[4], s[5] = c1, c2, c3
}

func (s *Side) Y1() (c1, c2, c3 Cell) {
	return s[3], s[4], s[5]
}

func (s *Side) setY2(c1, c2, c3 Cell) {
	s[6], s[7], s[8] = c1, c2, c3
}

func (s *Side) Y2() (c1, c2, c3 Cell) {
	return s[6], s[7], s[8]
}
