package day02

import (
	"bufio"
	"io"
)

func validateBlocks(blocks []block) bool {
	for _, block := range blocks {
		if block.Red > 12 || block.Green > 13 || block.Blue > 14 {
			return false
		}
	}
	return true
}

func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		d := s.Bytes()
		gameId, blocks := parseLine(d)
		if validateBlocks(blocks) {
			summe += gameId
		}
	}
	return summe, nil
}
