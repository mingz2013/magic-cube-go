package main

import (
	"log"
	"magic-cube-go/magic_cube"
)

func main() {

	m := magic_cube.NewMagicCube()
	//m.Random()

	al := magic_cube.RandomActionList()

	m.DoActionList(al)

	log.Println(m)

}
