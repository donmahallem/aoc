package day01

import (
	"bufio"
	"io"
)

func parseInputGen(in io.Reader) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		s := bufio.NewScanner(in)
		for s.Scan() {
			line := s.Bytes()
			if len(line) < 2 {
				continue
			}
			distance := 0
			for i := 1; i < len(line); i++ {
				distance = distance*10 + int(line[i]-'0')
			}
			switch line[0] {
			case 'R':
				// nothing to do
			case 'L':
				distance = -distance
			default:
				continue
			}
			if !yield(distance) {
				return
			}
		}
	}
}
