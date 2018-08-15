package magic_cube

import (
	"log"
	"math/rand"
	"time"
)

const (
	A_F  = "F"
	A_FR = "FR"
	A_F2 = "F2"

	A_B  = "B"
	A_BR = "BR"
	A_B2 = "B2"

	A_L  = "L"
	A_LR = "LR"
	A_L2 = "L2"

	A_R  = "R"
	A_RR = "RR"
	A_R2 = "R2"

	A_U  = "U"
	A_UR = "UR"
	A_U2 = "U2"

	A_D  = "D"
	A_DR = "DR"
	A_D2 = "D2"

	A_E  = "E"
	A_ER = "ER"
	A_E2 = "E2"

	A_S  = "S"
	A_SR = "SR"
	A_S2 = "S2"

	A_M  = "M"
	A_MR = "MR"
	A_M2 = "M2"
)

func JoinActionList(actionList []string) (s string) {
	for i := 0; i < len(actionList); i++ {
		action := actionList[i]

		if action == A_R {
			s += "T"
		} else if action == A_RR {
			s += "TR"
		} else if action == A_R2 {
			s += "T2"
		} else {
			s += action
		}

	}
	return
}

func ActionList() (actionList []string) {
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

func GenerateRandomActionList() (actionList []string) {
	// 生成随机action列表
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		j := r.Intn(len(ActionList()))
		actionList = append(actionList, ActionList()[j])
	}
	//log.Println(actionList)
	return
}

func ParseActionList(actions string) (actionList []string) {
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

func RandomActionList() []string {
	s := JoinActionList(GenerateRandomActionList())
	log.Println(s)
	al := ParseActionList(s)
	//log.Println(al)
	return al
}
