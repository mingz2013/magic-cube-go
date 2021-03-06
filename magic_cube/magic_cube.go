package magic_cube

import "log"

// 第一种思路，按每个面分，一个面9个cell，一共6个面

// 以左近下方为原点
// 向右为X坐标
// 向前为Y坐标
// 向上为Z坐标

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

// 前F后B左L右R上U下D
const (
	F = 0 // front
	B = 1 // back
	L = 2 // left
	R = 3 // right
	U = 4 // up
	D = 5 // down
)

// 读取不同面

func (m *MagicCube) FSide() *Side {
	return &m[F]
}

func (m *MagicCube) BSide() *Side {
	return &m[B]
}

func (m *MagicCube) LSide() *Side {
	return &m[L]
}

func (m *MagicCube) RSide() *Side {
	return &m[R]
}

func (m *MagicCube) USide() *Side {
	return &m[U]
}

func (m *MagicCube) DSide() *Side {
	return &m[D]
}

func (m *MagicCube) IsBingo() (b bool) {
	for i := 0; i < len(m); i++ {
		s := m[i]
		b = s.isBingo()
		if !b {
			return
		}
	}
	return true
}

func (m *MagicCube) DoActionList(actionList []string) {
	for i := 0; i < len(actionList); i++ {
		m.DoAction(actionList[i])
		b := m.IsBingo()
		if b {
			log.Println("isBingo.......")
			break
		}
	}
}

func (m *MagicCube) DoAction(action string) {
	log.Println("DoAction", action)
	switch action {
	case A_F:
		m.F()
	case A_FR:
		m.FR()
	case A_F2:
		m.F2()
	case A_B:
		m.B()
	case A_BR:
		m.BR()
	case A_B2:
		m.B2()
	case A_L:
		m.L()
	case A_LR:
		m.LR()
	case A_L2:
		m.L2()
	case A_R:
		m.R()
	case A_RR:
		m.RR()
	case A_R2:
		m.R2()
	case A_U:
		m.U()
	case A_UR:
		m.UR()
	case A_U2:
		m.U2()
	case A_D:
		m.D()
	case A_DR:
		m.DR()
	case A_D2:
		m.D2()
	case A_E:
		m.E()
	case A_ER:
		m.ER()
	case A_E2:
		m.E2()
	case A_S:
		m.S()
	case A_SR:
		m.SR()
	case A_S2:
		m.S2()
	case A_M:
		m.M()
	case A_MR:
		m.MR()
	case A_M2:
		m.M2()
	default:
		break
	}
}

// 一系列操作，旋转操作
func (m *MagicCube) U() {
	// Z2
	// 上层做顺时针90度旋转
	// 涉及前后左右四面的上层，涉及U的整个变换

	m.USide().clockwise()

	r := m.LSide().Y0()
	m.LSide().setY0(m.FSide().Y0())
	m.FSide().setY0(m.RSide().Y0())
	m.RSide().setY0(m.BSide().Y0())
	m.BSide().setY0(r)

}

func (m *MagicCube) UR() {
	// 上层做逆时针90度旋转

	m.USide().anticlockwise()

	r := m.LSide().Y0()
	m.LSide().setY0(m.BSide().Y0())
	m.BSide().setY0(m.RSide().Y0())
	m.RSide().setY0(m.FSide().Y0())
	m.FSide().setY0(r)
}

func (m *MagicCube) U2() {
	// 上层顺时针旋转180度，或，上层逆时针旋转180度

	m.U()
	m.U()

}

func (m *MagicCube) R() {
	// X2
	// 右侧

	m.RSide().clockwise()

	r := m.USide().X2()

	m.USide().setX2(m.FSide().X2())
	m.FSide().setX2(m.DSide().X2())
	m.DSide().setX2(m.BSide().X2())
	m.BSide().setX2(r)

}

func (m *MagicCube) RR() {
	m.RSide().anticlockwise()

	r := m.USide().X2()

	m.USide().setX2(m.BSide().X2())
	m.BSide().setX2(m.DSide().X2())
	m.DSide().setX2(m.FSide().X2())
	m.FSide().setX2(r)
}

func (m *MagicCube) R2() {
	m.R()
	m.R()
}

func (m *MagicCube) F() {
	// Y0
	// 前面

	//log.Println("<< F..")

	m.FSide().clockwise()

	r := m.USide().Y2()

	m.USide().setY2(m.LSide().X2())
	m.LSide().setX2(m.DSide().Y2())
	m.DSide().setY2(m.RSide().X0())
	m.RSide().setX0(r)

}
func (m *MagicCube) FR() {
	m.FSide().anticlockwise()

	r := m.USide().Y2()

	m.USide().setY2(m.RSide().X0())
	m.RSide().setX0(m.DSide().Y2())
	m.DSide().setY2(m.LSide().X2())
	m.LSide().setX2(r)
}

func (m *MagicCube) F2() {
	m.F()
	m.F()
}

func (m *MagicCube) L() {
	// X0

	m.LSide().clockwise()

	r := m.USide().X0()
	m.USide().setX0(m.BSide().X0())
	m.BSide().setX0(m.DSide().X0())
	m.DSide().setX0(m.FSide().X0())
	m.FSide().setX0(r)

}

func (m *MagicCube) LR() {
	m.LSide().anticlockwise()

	r := m.USide().X0()
	m.USide().setX0(m.FSide().X0())
	m.FSide().setX0(m.DSide().X0())
	m.DSide().setX0(m.BSide().X0())
	m.BSide().setX0(r)

}

func (m *MagicCube) L2() {
	m.L()
	m.L()
}

func (m *MagicCube) B() {
	// Y2
	m.BSide().clockwise()

	r := m.USide().Y0()
	m.USide().setY0(m.RSide().X2())
	m.RSide().setX2(m.DSide().Y0())
	m.DSide().setY0(m.LSide().X0())
	m.LSide().setX0(r)
}

func (m *MagicCube) BR() {
	m.BSide().anticlockwise()

	r := m.USide().Y0()
	m.USide().setY0(m.LSide().X0())
	m.LSide().setX0(m.DSide().Y0())
	m.DSide().setY0(m.RSide().X2())
	m.RSide().setX2(r)

}

func (m *MagicCube) B2() {
	m.R()
	m.R()
}

func (m *MagicCube) D() {
	// Z0

	m.DSide().clockwise()
	r := m.FSide().Y2()
	m.FSide().setY2(m.RSide().Y2())
	m.RSide().setY2(m.BSide().Y2())
	m.BSide().setY2(m.LSide().Y2())
	m.LSide().setY2(r)
}

func (m *MagicCube) DR() {
	m.DSide().anticlockwise()

	r := m.FSide().Y2()
	m.FSide().setY2(m.LSide().Y2())
	m.LSide().setY2(m.BSide().Y2())
	m.BSide().setY2(m.RSide().Y2())
	m.RSide().setY2(r)
}

func (m *MagicCube) D2() {
	m.D()
	m.D()
}

func (m *MagicCube) E() {
	// Z1
	// 中间层
	r := m.FSide().Y1()

	m.FSide().setY1(m.RSide().Y1())
	m.RSide().setY1(m.BSide().Y1())
	m.BSide().setY1(m.LSide().Y1())

	m.LSide().setY1(r)
}

func (m *MagicCube) ER() {
	r := m.FSide().Y1()

	m.FSide().setY1(m.LSide().Y1())
	m.LSide().setY1(m.BSide().Y1())
	m.BSide().setY1(m.RSide().Y1())

	m.RSide().setY1(r)
}

func (m *MagicCube) E2() {
	m.E()
	m.E()
}

func (m *MagicCube) M() {
	// X1

	r := m.FSide().X1()
	m.FSide().setX1(m.USide().X1())
	m.USide().setX1(m.BSide().X1())
	m.BSide().setX1(m.DSide().X1())
	m.DSide().setX1(r)

}

func (m *MagicCube) MR() {
	r := m.FSide().X1()
	m.FSide().setX1(m.DSide().X1())
	m.DSide().setX1(m.BSide().X1())
	m.BSide().setX1(m.USide().X1())
	m.USide().setX1(r)
}

func (m *MagicCube) M2() {
	m.M()
	m.M()

}

func (m *MagicCube) S() {
	// Y1
	r := m.USide().Y1()
	m.USide().setY1(m.LSide().X1())
	m.LSide().setX1(m.DSide().Y1())
	m.DSide().setY1(m.RSide().X1())
	m.RSide().setX1(r)
}

func (m *MagicCube) SR() {
	r := m.USide().Y1()
	m.USide().setY1(m.RSide().X1())
	m.RSide().setX1(m.DSide().Y1())
	m.DSide().setY1(m.LSide().X1())
	m.LSide().setX1(r)
}

func (m *MagicCube) S2() {
	m.S()
	m.S()
}
