package day08

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	games, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	steps := 0
	numInstructions := len(games.Instructions)
	current := games.Start
	for ; ; steps++ {
		if current == games.End {
			break
		}
		if games.Instructions[steps%numInstructions] {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return steps, nil
}
