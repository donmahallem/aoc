package day22

import (
	"io"
	"slices"
)

type supportMap = [][]int

func settleBricks(bricks []brick) {
	slices.SortFunc(bricks, func(a, b brick) int {
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

func buildSupportGraphMap(bricks []brick) (supportMap, supportMap) {
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

func countRedundantBricks(bricks []brick) (total int) {
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

func Part1(in io.Reader) (int, error) {
	inp, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	settleBricks(inp)
	return countRedundantBricks(inp), nil
}
