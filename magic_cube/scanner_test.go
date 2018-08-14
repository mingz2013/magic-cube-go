package magic_cube

import (
	"log"
	"testing"
)

func TestScanner_NextToken(t *testing.T) {
	s := NewScanner("LBRFTRLBRTRT2L2B2BBBLLDUURDRDDRLLL")

	for {
		tok, ok := s.NextToken()
		log.Println(tok, ok)
		if !ok {
			log.Println("break...")
			break
		}

	}
}
