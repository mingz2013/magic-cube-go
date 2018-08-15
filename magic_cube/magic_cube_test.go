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

	log.Println(m)

}
