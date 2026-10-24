package day08

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part1(in io.Reader) (int, error) {
	games, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	if games.Start == nil || games.End == nil {
		return 0, aoc_utils.NewParseError("start node AAA or end node ZZZ is missing", nil)
	}
	steps := 0
	numInstructions := len(games.Instructions)
	// every (node, instruction index) state repeats after this many steps
	maxSteps := len(games.Nodes) * numInstructions
	current := games.Start
	for ; ; steps++ {
		if current == games.End {
			break
		}
		if steps > maxSteps {
			return 0, aoc_utils.NewParseError("end node ZZZ is unreachable", nil)
		}
		if games.Instructions[steps%numInstructions] {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return steps, nil
}
