package day03

import (
	"bufio"
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

// keepRule to pass in evaluation rules
type keepRule func(ones, zeros int) bool

// oxygenRule keeps the most common bit, preferring '1' on ties.
func oxygenRule(ones, zeros int) bool {
	return ones >= zeros
}

// co2Rule keeps the least common bit, preferring '0' on ties and when only '0' is present.
func co2Rule(ones, zeros int) bool {
	return !(zeros > 0 && (ones == 0 || zeros <= ones))
}

func findRating(candidates []int, bitLen int, keepOne keepRule) int {
	for bit := bitLen - 1; bit >= 0 && len(candidates) > 1; bit-- {
		ones := 0
		for _, v := range candidates {
			ones += v >> bit & 1
		}
		want := 0
		if keepOne(ones, len(candidates)-ones) {
			want = 1
		}
		n := 0
		for _, v := range candidates {
			if v>>bit&1 == want {
				candidates[n] = v
				n++
			}
		}
		candidates = candidates[:n]
	}
	return candidates[0]
}

func Part2(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	expectedInputLength := 0
	values := make([]int, 0, 1024)

	for s.Scan() {
		line := s.Bytes()
		if expectedInputLength == 0 {
			expectedInputLength = len(line)
		} else if len(line) != expectedInputLength {
			return 0, aoc_utils.NewParseErrorf("unexpected line length: got %d, expected %d", len(line), expectedInputLength)
		}
		v := 0
		for _, b := range line {
			v <<= 1
			if b == '1' {
				v |= 1
			}
		}
		values = append(values, v)
	}

	if err := s.Err(); err != nil {
		return 0, err
	}
	if len(values) == 0 {
		return 0, aoc_utils.NewParseErrorf("empty input")
	}

	co2Candidates := slices.Clone(values)
	oxygen := findRating(values, expectedInputLength, oxygenRule)
	co2 := findRating(co2Candidates, expectedInputLength, co2Rule)

	return oxygen * co2, nil
}
