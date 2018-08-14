package magic_cube

import "log"

type Scanner struct {
	src string

	offset int

	ch rune
}

func (s *Scanner) Init(src string) {
	s.src = src
	s.offset = 0
	s.ch = rune(s.src[s.offset])
}

func NewScanner(src string) Scanner {
	s := Scanner{}
	s.Init(src)
	return s
}

func (s *Scanner) NextToken() (tok string, ok bool) {

	return s.scanToken()
}

func (s *Scanner) nextCh() {
	s.offset += 1

	if s.offset >= len(s.src) {
		s.ch = -1

	} else {
		s.ch = rune(s.src[s.offset])
	}

}

func (s *Scanner) scanToken() (tok string, ok bool) {
	log.Println("scanToken...", tok, string(s.ch), s.offset)

	switch s.ch {
	case 'F':
		tok = "F"
	case 'B':
		tok = "B"
	case 'U':
		tok = "U"
	case 'D':
		tok = "D"
	case 'L':
		tok = "L"
	case 'T':
		tok = "R"
	case 'E':
		tok = "E"
	case 'M':
		tok = "M"
	case 'S':
		tok = "S"
	default:
		tok = ""
		ok = false
		return
	}

	s.nextCh()
	tok, ok = s.scanAfter(tok)

	return
}

func (s *Scanner) scanAfter(tok string) (string, bool) {

	log.Println("scanAfter...", tok, s.ch, s.offset)

	switch s.ch {
	case 'R':
		tok += "R"
	case '2':
		tok += "2"
	default:

		return tok, true
	}

	s.nextCh()
	return tok, true

}
