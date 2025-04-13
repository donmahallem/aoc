package day19

import (
	"bufio"
	"io"
	"strings"
)

type TowelPattern = string
type TowelPatterns = map[TowelPattern]uint
type Towel = string
type Towels = []Towel

func ParseFirstLine(line *string) (*TowelPatterns, *uint) {
	numbers := strings.Split((*line), ",")
	patterns := make(TowelPatterns, len(numbers))
	var longestPattern uint = 0
	var key string
	for _, num := range numbers {
		key = strings.Trim(num, " ")
		patterns[key] = 1
		if keyLen := uint(len(key)); keyLen > longestPattern {
			longestPattern = keyLen
		}
	}
	return &patterns, &longestPattern
}

func ParseInput(in io.Reader) (*TowelPatterns, *Towels, *uint) {
	var patterns *TowelPatterns
	var keyLen *uint
	towels := make(Towels, 0)
	firstLine := true
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			firstLine = false
			continue
		} else if firstLine {
			patterns, keyLen = ParseFirstLine(&line)
		} else {
			towels = append(towels, line)
		}
	}
	return patterns, &towels, keyLen
}

func Gogo(patterns *TowelPatterns, towel *Towel, toCheck *[]Towel) bool {
	//toCheck := make([]Towel, 0)
	*toCheck = append((*toCheck)[:0], *towel)
	var currentTowel Towel
	for {
		currentStackLength := len((*toCheck))
		if currentStackLength == 0 {
			break
		}
		currentTowel = (*toCheck)[currentStackLength-1]
		(*toCheck) = (*toCheck)[0 : currentStackLength-1]
		for pattern := range *patterns {
			if pattern == currentTowel {
				return true
			} else if strings.HasPrefix(currentTowel, pattern) {
				(*toCheck) = append((*toCheck), currentTowel[len(pattern):])
			}
		}
	}
	return false
}

func Part1(in io.Reader) uint {
	patterns, towls, _ := ParseInput(in)
	toCheck := make([]Towel, 0)
	var count uint = 0
	for _, towl := range *towls {
		if Gogo(patterns, &towl, &toCheck) {
			count++
		}
	}
	return count
}
