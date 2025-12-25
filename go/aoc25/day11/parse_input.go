package day11

import (
	"bufio"
	"io"

	aoc_utils "github.com/donmahallem/aoc/go/aoc_utils"
)

func parseInput(in io.Reader) (map[uint64][]uint64, error) {
	s := bufio.NewScanner(in)

	nodeMap := make(map[uint64][]uint64)
	for s.Scan() {
		line := s.Bytes()

		sourceKey := uint64(0)
		currentKey := uint64(0)
		inSourceKey := true
		targets := make([]uint64, 0, 4)
		for _, c := range line {
			if c == ':' {
				inSourceKey = false
				sourceKey = currentKey
				currentKey = 0
			} else if c >= 'a' && c <= 'z' {
				currentKey = currentKey<<8 + uint64(c)
			} else if c == ' ' {
				if !inSourceKey && currentKey != 0 {
					targets = append(targets, currentKey)
					currentKey = 0
				}
			} else {
				return nil, aoc_utils.NewUnexpectedInputError(c)
			}
			// Any other character (like \r) is simply ignored, preventing key corruption
		}
		if currentKey != 0 {
			targets = append(targets, currentKey)
		}
		nodeMap[sourceKey] = targets
	}
	return nodeMap, nil
}
