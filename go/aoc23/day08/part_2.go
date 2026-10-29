package day08

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func getCycleSize(instructions pathInstructions, n *node, maxSteps int) (uint, error) {
	current := n
	for i := 0; i <= maxSteps; i++ {
		if current == nil {
			return 0, aoc_utils.NewParseError("Reached nil node in cycle detection", nil)
		}
		if i > 0 && current.EndsInZ && i%len(instructions) == 0 {
			return uint(i), nil
		}
		if instructions[i%len(instructions)] {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return 0, aoc_utils.NewSolverErrorf("no node ending in Z reached after %d steps", maxSteps)
}

func Part2(in io.Reader) (uint, error) {
	games, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	var currentBase uint = 1
	maxSteps := len(games.Nodes) * len(games.Instructions)
	for _, n := range games.Nodes {
		if n.EndsInA {
			cycleSize, err := getCycleSize(games.Instructions, n, maxSteps)
			if err != nil {
				return 0, err
			}
			currentBase = int_util.LcmInt(currentBase, cycleSize)
		}
	}
	return currentBase, nil
}
