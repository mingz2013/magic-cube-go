package magic_cube

import (
	"log"
	"math/rand"
	"time"
)

// 第一种思路，按每个面分，一个面9个cell，一共6个面

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

func (m *MagicCube) Random() {
	m.doActionList(m.generateRandomActionList())
}

func (m *MagicCube) actionList() (actionList []string) {
	return []string{
		A_F, A_FR, A_F2,
		A_B, A_BR, A_B2,
		A_L, A_LR, A_L2,
		A_R, A_RR, A_R2,
		A_U, A_UR, A_U2,
		A_D, A_DR, A_D2,
		A_E, A_ER, A_E2,
		A_S, A_SR, A_S2,
		A_M, A_MR, A_M2,
	}
}

func (m *MagicCube) generateRandomActionList() (actionList []string) {
	// 生成随机action列表
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		j := r.Intn(len(m.actionList()))
		actionList = append(actionList, m.actionList()[j])
	}
	log.Println(actionList)
	return
}

func (m *MagicCube) parseActionList(actions string) (actionList []string) {
	// 拆分string为一个个单独的action

	s := NewScanner(actions)

	for {
		tok, ok := s.NextToken()
		if !ok {
			break
		}

		actionList = append(actionList, tok)
	}

	return
}

func (m *MagicCube) doActionList(actionList []string) {
	for i := 0; i < len(actionList); i++ {
		m.DoAction(actionList[i])
	}
}

func (m *MagicCube) DoAction(action string) {
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
	// 上层做顺时针90度旋转
	// 涉及前后左右四面的上层，涉及U的整个变换

	m.USide().clockwise()

	l1, l2, l3 := m.LSide().Y0()
	m.LSide().changeY0(m.FSide().Y0())
	m.FSide().changeY0(m.RSide().Y0())
	m.RSide().changeY0(m.BSide().Y0())
	m.BSide().changeY0(l1, l2, l3)

}

func (m *MagicCube) UR() {
	// 上层做逆时针90度旋转

	m.USide().anticlockwise()

	l1, l2, l3 := m.LSide().Y0()
	m.LSide().changeY0(m.BSide().Y0())
	m.BSide().changeY0(m.RSide().Y0())
	m.RSide().changeY0(m.RSide().Y0())
	m.RSide().changeY0(l1, l2, l3)
}

func (m *MagicCube) U2() {
	// 上层顺时针旋转180度，或，上层逆时针旋转180度

	m.U()
	m.U()

}

func (m *MagicCube) R() {
	// 右侧

	m.RSide().clockwise()

	l1, l2, l3 := m.USide().X2()

	m.USide().changeX2(m.FSide().X2())
	m.FSide().changeX2(m.DSide().X2())
	m.DSide().changeX2(m.BSide().X2())
	m.BSide().changeX2(l1, l2, l3)

}

func (m *MagicCube) RR() {
	m.RSide().anticlockwise()

	l1, l2, l3 := m.USide().X2()

	m.USide().changeX2(m.BSide().X2())
	m.BSide().changeX2(m.DSide().X2())
	m.DSide().changeX2(m.FSide().X2())
	m.FSide().changeX2(l1, l2, l3)
}

func (m *MagicCube) R2() {
	m.R()
	m.R()
}

func (m *MagicCube) F() {
	// 前面

	//log.Println("<< F..")

	m.FSide().clockwise()

	l1, l2, l3 := m.USide().Y2()

	m.USide().changeY2(m.LSide().X2())
	m.LSide().changeX2(m.DSide().Y2())
	m.DSide().changeY2(m.RSide().X0())
	m.RSide().changeX0(l1, l2, l3)

}
func (m *MagicCube) FR() {
	m.FSide().anticlockwise()

	l1, l2, l3 := m.USide().Y2()

	m.USide().changeY2(m.RSide().X0())
	m.RSide().changeX0(m.DSide().Y2())
	m.DSide().changeY2(m.LSide().X2())
	m.LSide().changeX2(l1, l2, l3)
}

func (m *MagicCube) F2() {
	m.F()
	m.F()
}

func (m *MagicCube) L() {

	m.LSide().clockwise()

	l1, l2, l3 := m.USide().X0()
	m.USide().changeX0(m.BSide().Y2())
	m.BSide().changeY2(m.DSide().X0())
	m.DSide().changeX0(m.FSide().X0())
	m.FSide().changeX0(l1, l2, l3)

}

func (m *MagicCube) LR() {
	m.LSide().anticlockwise()

	l1, l2, l3 := m.USide().X0()
	m.USide().changeX0(m.FSide().X0())
	m.FSide().changeX0(m.DSide().X0())
	m.DSide().changeX0(m.LSide().Y2())
	m.LSide().changeY2(l1, l2, l3)

}

func (m *MagicCube) L2() {
	m.L()
	m.L()
}

func (m *MagicCube) B() {
	m.BSide().clockwise()

	l1, l2, l3 := m.USide().Y0()
	m.USide().changeY0(m.RSide().X2())
	m.RSide().changeX2(m.DSide().Y0())
	m.DSide().changeY0(m.LSide().X0())
	m.LSide().changeX0(l1, l2, l3)
}

func (m *MagicCube) BR() {
	m.BSide().anticlockwise()

	l1, l2, l3 := m.USide().Y0()
	m.USide().changeY0(m.LSide().X0())
	m.LSide().changeX0(m.DSide().Y0())
	m.DSide().changeY0(m.RSide().X2())
	m.RSide().changeX2(l1, l2, l3)

}

func (m *MagicCube) B2() {
	m.R()
	m.R()
}

func (m *MagicCube) D() {
	m.DSide().clockwise()
	l1, l2, l3 := m.FSide().Y2()
	m.FSide().changeY2(m.RSide().Y2())
	m.RSide().changeY2(m.BSide().Y2())
	m.BSide().changeY2(m.LSide().Y2())
	m.LSide().changeY2(l1, l2, l3)
}

func (m *MagicCube) DR() {
	m.DSide().clockwise()
	l1, l2, l3 := m.FSide().Y2()
	m.FSide().changeY2(m.LSide().Y2())
	m.LSide().changeY2(m.BSide().Y2())
	m.BSide().changeY2(m.RSide().Y2())
	m.RSide().changeY2(l1, l2, l3)
}

func (m *MagicCube) D2() {
	m.D()
	m.D()
}

func (m *MagicCube) E() {
	// 中间层
}

func (m *MagicCube) ER() {

}

func (m *MagicCube) E2() {

}

func (m *MagicCube) S() {

}

func (m *MagicCube) SR() {

}

func (m *MagicCube) S2() {

}

func (m *MagicCube) M() {

}

func (m *MagicCube) MR() {

}

func (m *MagicCube) M2() {

}
