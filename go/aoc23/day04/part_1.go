package day04

import (
	"bufio"
	"io"
	"slices"
)

func countWinnings(winners []uint8, picks []uint8) int {
	score := 0
	for _, winner := range winners {
		if slices.Contains(picks, winner) {
			score++
		}
	}
	return score
}

func getScore(winners []uint8, picks []uint8) int {
	score := countWinnings(winners, picks)
	if score > 0 {
		return 1 << (score - 1)
	}
	return 0
}

func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	score := 0
	for s.Scan() {
		_, a, b := parseLine(s.Bytes())
		score += getScore(a, b)
	}
	return score, nil
}
