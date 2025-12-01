package day01

import (
	"bufio"
	"io"
)

type inputData struct {
	left     bool
	distance int
}

func parseInputGen(in io.Reader) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		s := bufio.NewScanner(in)
		for s.Scan() {
			line := s.Bytes()
			distance := 0
			for i := 1; i < len(line); i++ {
				distance = distance*10 + int(line[i]-'0')
			}
			if line[0] == 'L' {
				distance = -distance
			}
			if !yield(distance) {
				return
			}
		}
	}
}
