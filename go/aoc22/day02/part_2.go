package day02

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

// X lose
// Y draw
// Z win

func Part2(in io.Reader) (int, error) {
	parsedNumbers := make([]int, 0, 256)
	s := bufio.NewScanner(in)
	currentSum := 0
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}
		if len(line) < 3 {
			return 0, aoc_utils.NewParseErrorf("line '%s' is to short", line)
		}
		t1, err := mapToken(line[0])
		if err != nil {
			return 0, err
		}
		var t2 int
		switch line[2] {
		case 'X': // lose
			t2 = (t1 + 2) % 3
			if t2 == 0 {
				t2 = 3
			}
		case 'Y': // draw
			t2 = t1
		case 'Z': // win
			t2 = t1%3 + 1
		default:
			return 0, aoc_utils.NewUnexpectedInputError(line[2])
		}
		currentSum += t2
		switch line[2] {
		case 'X': // lose
			// no additional points
		case 'Y': // draw
			currentSum += 3
		case 'Z': // win
			currentSum += 6
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	if currentSum > 0 {
		parsedNumbers = append(parsedNumbers, currentSum)
	}
	return currentSum, nil
}
