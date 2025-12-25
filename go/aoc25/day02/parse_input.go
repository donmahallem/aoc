package day02

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/math"
)

type intInterval = math.Interval[uint64]

func parseInputGen(in io.Reader) func(yield func(intInterval) bool) {
	// cap parsed values to avoid excessive work from fuzzed inputs; set high
	// enough to preserve real sample/test values.
	const maxIntervalValue uint64 = 1_000_000_000_000
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
						digit := uint64(c - '0')
						// clamp to avoid huge numeric values that make the solver run
						// for very long or allocate massive structures.
						if *currentNum > maxIntervalValue/10 || (*currentNum)*10+digit > maxIntervalValue {
							*currentNum = maxIntervalValue
						} else {
							*currentNum = (*currentNum)*10 + digit
						}
					}
				}
			}
		}
		yield(inp)
	}
}
