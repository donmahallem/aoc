package day10

import (
	"bufio"
	"io"
	"slices"
)

func Part2(in io.Reader) (int, error) {
	scanner := bufio.NewScanner(in)
	scores := make([]int, 0, 128)
	for scanner.Scan() {
		line := scanner.Bytes()
		score, ast, err := parseAst(line)
		if score > 0 {
			continue
		} else if err != nil {
			return 0, err
		}
		score = 0
		for i := len(ast) - 1; i >= 0; i-- {
			score *= 5
			switch ast[i] {
			case '(':
				score += 1
			case '[':
				score += 2
			case '{':
				score += 3
			case '<':
				score += 4
			}
		}
		scores = append(scores, score)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	slices.Sort(scores)
	return scores[len(scores)/2], nil
}
