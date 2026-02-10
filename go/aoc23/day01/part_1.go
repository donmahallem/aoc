package day01

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

func parseLinePart1(line []byte) (int, error) {

	var searchNumber func(startIdx, endIdx, direction int) (int, int)
	searchNumber = func(startIdx, endIdx, direction int) (idx int, res int) {
		for i := startIdx; i != endIdx; i += direction {
			if val, ok := bytes.ParseIntFromByte[int](line[i]); ok {
				return i, val
			}
		}
		return -1, 0
	}
	idxLeft, valLeft := searchNumber(0, len(line), 1)
	if idxLeft < 0 {
		return 0, aoc_utils.NewParseError("No number found", nil)
	}
	idxRight, valRight := searchNumber(len(line)-1, idxLeft, -1)
	if idxRight < 0 {
		// only one number found
		return valLeft*10 + valLeft, nil
	}
	return valLeft*10 + valRight, nil
}
func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		val, err := parseLinePart1(s.Bytes())
		if err != nil {
			return 0, err
		}
		summe += val
	}
	return summe, nil
}
