package day20

import (
	"io"
)

func CountCheats2(racewayPoints *[]Point, minSavings int, maxCheatDistance int) int {
	cheatCount := 0
	for leftIdx := range len(*racewayPoints) - minSavings {
		for rightIdx := len(*racewayPoints) - 1; leftIdx+minSavings < rightIdx; rightIdx-- {
			dst := (*racewayPoints)[leftIdx].DistanceManhatten((*racewayPoints)[rightIdx])
			if dst < 2 || dst > maxCheatDistance {
				continue
			}
			savings := rightIdx - leftIdx - dst
			if savings >= minSavings {
				cheatCount++
			}
		}
	}
	return cheatCount
}

func Part2(in io.Reader) int {
	patterns := ParseInput(in)
	return CountCheats2(patterns, 100, 20)
}
