package day22

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type brick struct {
	X1, Y1, Z1, X2, Y2, Z2 int
}

func (b *brick) IntersectXY(other *brick) bool {
	return !(b.X2 < other.X1 || b.X1 > other.X2 || b.Y2 < other.Y1 || b.Y1 > other.Y2)
}
func parseInput(r io.Reader) (bricks []brick, err error) {

	scanner := bufio.NewScanner(r)

	parseBrick := func(line []byte) (*brick, error) {
		var b brick
		parseOrder := [6]*int{&b.X1, &b.Y1, &b.Z1, &b.X2, &b.Y2, &b.Z2}
		currentOrderIndex := 0
		currentValue := 0
		for i, c := range line {
			if c >= '0' && c <= '9' {
				currentValue = currentValue*10 + int(c-'0')
			} else if c == ',' || c == '~' || i == len(line)-1 {
				if currentOrderIndex >= len(parseOrder) {
					return nil, aoc_utils.NewParseError("too many values in brick definition", nil)
				}
				*parseOrder[currentOrderIndex] = currentValue
				currentOrderIndex++
				currentValue = 0
			}
		}
		if currentValue > 0 && currentOrderIndex < len(parseOrder) {
			*parseOrder[currentOrderIndex] = currentValue
		}
		return &b, nil
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		brick, err := parseBrick(line)
		if err != nil {
			return nil, err
		}
		bricks = append(bricks, *brick)
	}
	return
}
