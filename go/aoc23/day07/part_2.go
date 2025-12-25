package day07

import (
	"io"
	"slices"
)

func Part2(in io.Reader) (int, error) {
	games, err := parseInput2(in)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(games, func(a game, b game) int {
		return a.HandHash - b.HandHash
	})
	var total int = 0
	for idx, game := range games {
		total += game.Bid * int(idx+1)
	}
	return total, nil
}
