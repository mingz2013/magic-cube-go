package magic_cube

import (
	"log"
	"testing"
)

func TestMagicCube_IsBingo(t *testing.T) {

	m := NewMagicCube()

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("UUR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("TTR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("FFR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("LLR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("BBR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("DDR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("EER"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("MMR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("SSR"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("U2U2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("T2T2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("F2F2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("L2L2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("B2B2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("D2D2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("E2E2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("M2M2"))

	log.Println(m.IsBingo())

	m.DoActionList(ParseActionList("S2S2"))

	log.Println(m.IsBingo())

	log.Println(m)

}
