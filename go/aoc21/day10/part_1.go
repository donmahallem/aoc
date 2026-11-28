package day10

import (
	"bufio"
	"io"
)

func Part1(in io.Reader) (int, error) {
	scanner := bufio.NewScanner(in)
	count := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		score, _, err := parseAst(line)
		if score > 0 {
			count += score
		} else if err != nil {
			return 0, err
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}
