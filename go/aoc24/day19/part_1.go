package day19

import (
	"bufio"
	"io"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type TowelPattern = string
type TowelPatterns = map[TowelPattern]uint
type Towel = string
type Towels = []Towel

func ParseFirstLine(line string) (TowelPatterns, uint, error) {
	numbers := strings.Split(line, ",")
	patterns := make(TowelPatterns, len(numbers))
	var longestPattern uint = 0
	var key string
	for _, num := range numbers {
		key = strings.Trim(num, " ")
		if len(key) == 0 {
			return nil, 0, aoc_utils.NewUnexpectedInputError(0)
		}
		patterns[key] = 1
		if keyLen := uint(len(key)); keyLen > longestPattern {
			longestPattern = keyLen
		}
	}
	return patterns, longestPattern, nil
}

func ParseInput(in io.Reader) (TowelPatterns, Towels, uint, error) {
	var patterns TowelPatterns
	var keyLen uint
	towels := make(Towels, 0)
	firstLine := true
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			firstLine = false
			continue
		} else if firstLine {
			var err error
			patterns, keyLen, err = ParseFirstLine(line)
			if err != nil {
				return nil, nil, 0, err
			}
		} else {
			towels = append(towels, line)
		}
	}
	return patterns, towels, keyLen, nil
}

func Gogo(patterns TowelPatterns, towel Towel, toCheck *[]Towel) bool {
	*toCheck = append((*toCheck)[:0], towel)
	visited := make(map[Towel]struct{})
	visited[towel] = struct{}{}
	for {
		currentStackLength := len(*toCheck)
		if currentStackLength == 0 {
			break
		}
		currentTowel := (*toCheck)[currentStackLength-1]
		(*toCheck) = (*toCheck)[0 : currentStackLength-1]
		for pattern := range patterns {
			if pattern == currentTowel {
				return true
			} else if strings.HasPrefix(currentTowel, pattern) {
				next := currentTowel[len(pattern):]
				if _, ok := visited[next]; !ok {
					visited[next] = struct{}{}
					(*toCheck) = append((*toCheck), next)
				}
			}
		}
	}
	return false
}

func Part1(in io.Reader) (uint, error) {
	patterns, towls, _, err := ParseInput(in)
	if err != nil {
		return 0, err
	}
	toCheck := make([]Towel, 0)
	var count uint = 0
	for _, towl := range towls {
		if Gogo(patterns, towl, &toCheck) {
			count++
		}
	}
	return count, nil
}
