package day22

import (
	"io"
)

// counts bricks that are only supported by the given brick idx
func countSupportedBricks(bricks []Brick, supports, supportedBy supportMap) int {
	fallen := make([]bool, len(supports))
	queue := make([]int, 0, len(supports))

	total := 0
	for i := range bricks {
		fallenCount := 0
		queue = append(queue, i)
		fallen[i] = true
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			for _, above := range supports[current] {
				if fallen[above] {
					continue
				}
				shouldFall := true
				for _, supporter := range supportedBy[above] {
					if !fallen[supporter] {
						shouldFall = false
						break
					}
				}
				if shouldFall {
					fallen[above] = true
					fallenCount++
					queue = append(queue, above)
				}
			}
		}
		total += fallenCount
		fallenCount = 0
		clear(fallen)
	}
	return total
}
func Part2(in io.Reader) int {
	bricks := ParseInput(in)
	settleBricks(bricks)
	supportedBy, supports := buildSupportGraphMap(bricks)

	return countSupportedBricks(bricks, supports, supportedBy)
}
