package day22

import (
	"bufio"
	"io"
	"slices"
)

type supportMap = [][]int

type Brick struct {
	X1, Y1, Z1, X2, Y2, Z2 int
}

func (b Brick) IntersectXY(other *Brick) bool {
	return !(b.X2 < other.X1 || b.X1 > other.X2 || b.Y2 < other.Y1 || b.Y1 > other.Y2)
}

func ParseInput(r io.Reader) (bricks []Brick) {

	scanner := bufio.NewScanner(r)

	parseBrick := func(line []byte) Brick {
		var b Brick
		parseOrder := [6]*int{&b.X1, &b.Y1, &b.Z1, &b.X2, &b.Y2, &b.Z2}
		currentOrderIndex := 0
		currentValue := 0
		for i, c := range line {
			if c >= '0' && c <= '9' {
				currentValue = currentValue*10 + int(c-'0')
			} else if c == ',' || c == '~' || i == len(line)-1 {
				*parseOrder[currentOrderIndex] = currentValue
				currentOrderIndex++
				currentValue = 0
			}
		}
		if currentValue > 0 && currentOrderIndex < len(parseOrder) {
			*parseOrder[currentOrderIndex] = currentValue
		}
		return b
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		bricks = append(bricks, parseBrick(line))
	}
	return
}

func settleBricks(bricks []Brick) {
	slices.SortFunc(bricks, func(a, b Brick) int {
		if a.Z1 != b.Z1 {
			return a.Z1 - b.Z1
		} else if a.X1 != b.X1 {
			return a.X1 - b.X1
		} else if a.Y1 != b.Y1 {
			return a.Y1 - b.Y1
		} else {
			return 0
		}
	})

	for currentIdx := range bricks {
		currentBrick := &bricks[currentIdx]

		maxSupport := 0
		for checkIdx := currentIdx - 1; checkIdx >= 0; checkIdx-- {
			checkBrick := &bricks[checkIdx]
			if !currentBrick.IntersectXY(checkBrick) {
				continue
			}
			if checkBrick.Z2 > maxSupport {
				maxSupport = checkBrick.Z2
			}
		}

		height := currentBrick.Z2 - currentBrick.Z1
		newZ1 := maxSupport + 1
		currentBrick.Z1 = newZ1
		currentBrick.Z2 = newZ1 + height
	}

}

func buildSupportGraphMap(bricks []Brick) (supportMap, supportMap) {
	n := len(bricks)
	supportedBy := make(supportMap, n) // supporters of each brick
	supports := make(supportMap, n)    // bricks each brick supports

	for upper := range n {
		for lower := range upper {
			if !bricks[upper].IntersectXY(&bricks[lower]) {
				continue
			}
			if bricks[lower].Z2 == bricks[upper].Z1-1 {
				supportedBy[upper] = append(supportedBy[upper], lower)
				supports[lower] = append(supports[lower], upper)
			}
		}
	}
	return supportedBy, supports
}

func countRedundantBricks(bricks []Brick) (total int) {
	supportedBy, supports := buildSupportGraphMap(bricks)

	for i := range bricks {
		redundant := true
		for _, above := range supports[i] {
			if len(supportedBy[above]) == 1 {
				redundant = false
				break
			}
		}
		if redundant {
			total++
		}
	}
	return
}

func Part1(in io.Reader) int {
	inp := ParseInput(in)
	settleBricks(inp)
	return countRedundantBricks(inp)
}
