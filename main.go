package main

import (
	"github.com/mingz2013/magic-cube-go/magic_cube"
	"log"
)

func main() {

	m := magic_cube.NewMagicCube()
	//m.Random()

	al := magic_cube.RandomActionList()

	m.DoActionList(al)

	log.Println(m)

}
