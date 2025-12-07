package day02

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/math"
)

type intInterval = math.Interval[uint64]

func parseInputGen(in io.Reader) func(yield func(intInterval) bool) {
	return func(yield func(intInterval) bool) {
		s := bufio.NewScanner(in)
		inp := intInterval{}
		currentNum := &inp.Min
		for s.Scan() {
			line := s.Bytes()

			for i := range line {
				c := line[i]
				switch c {
				case '-':
					currentNum = &inp.Max
				case ',':
					if !yield(inp) {
						return
					}
					// reset
					currentNum = &inp.Min
					inp.Min = 0
					inp.Max = 0

				default:
					if c >= '0' && c <= '9' {
						*currentNum = (*currentNum)*10 + uint64(c-'0')
					}
				}
			}
		}
		yield(inp)
	}
}
