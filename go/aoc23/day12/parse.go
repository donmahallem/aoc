package day12

import (
	"bufio"
	"bytes"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type springState uint8

const (
	stateUnknown springState = '?'
	stateOk      springState = '.'
	stateBroken  springState = '#'
)

type line struct {
	BrokenParts []int
	States      []springState
}

/**
 * Parses a single line of input into a Line struct.
 * observationMultiplier determines how many times the observation is repeated.
 */
func parseLine(l []byte, observationMultiplier int) (*line, error) {
	splitIndex := bytes.IndexByte(l, ' ')
	switch splitIndex {
	case -1:
		return nil, aoc_utils.NewParseError("Unexpected input format", nil)
	case 0:
		return nil, aoc_utils.NewParseError("No spring states found", nil)
	}

	brokenPartSeries := bytes.Split(l[splitIndex+1:], []byte{','})
	numberOfBrokenGroups := len(brokenPartSeries)

	// sanity limits to avoid resource exhaustion
	const maxBrokenGroups = 1024
	const maxStatesBase = 10000
	if numberOfBrokenGroups == 0 || numberOfBrokenGroups > maxBrokenGroups {
		return nil, aoc_utils.NewParseError("invalid broken parts list", nil)
	}

	// parse broken part numbers safely (only decimal digits allowed)
	brokenParts := make([]int, numberOfBrokenGroups*observationMultiplier)
	for i := 0; i < numberOfBrokenGroups; i++ {
		token := bytes.TrimSpace(brokenPartSeries[i])
		if len(token) == 0 {
			return nil, aoc_utils.NewParseError("empty broken part token", nil)
		}
		num := 0
		for _, ch := range token {
			if ch < '0' || ch > '9' {
				return nil, aoc_utils.NewParseError("invalid broken part number", nil)
			}
			digit := int(ch - '0')
			// check for overflow / unreasonable counts
			if num > (1<<30-1-digit)/10 {
				return nil, aoc_utils.NewParseError("broken part number too large", nil)
			}
			num = num*10 + digit
		}
		// number of broken blocks in group cannot exceed the line length
		if num < 0 {
			return nil, aoc_utils.NewParseError("invalid broken part number", nil)
		}
		brokenParts[i] = num
	}

	// validate splitIndex (number of states)
	if splitIndex <= 0 || splitIndex > maxStatesBase {
		return nil, aoc_utils.NewParseError("invalid spring states length", nil)
	}
	// total states after observation multiplier
	statesLen := (splitIndex * observationMultiplier) + (observationMultiplier - 1)
	if statesLen <= 0 || statesLen > maxStatesBase*observationMultiplier {
		return nil, aoc_utils.NewParseError("resulting states length is unreasonable", nil)
	}

	states := make([]springState, statesLen)
	for idx, c := range l[:splitIndex] {
		switch c {
		case '?':
			states[idx] = stateUnknown
		case '#':
			states[idx] = stateBroken
		case '.':
			states[idx] = stateOk
		default:
			return nil, aoc_utils.NewUnexpectedInputError(c)
		}
	}
	if observationMultiplier > 1 {
		for idx := 0; idx < observationMultiplier-1; idx++ {
			offset := (splitIndex + 1) * (idx + 1)
			if offset-1 < 0 || offset+splitIndex > len(states) {
				return nil, aoc_utils.NewParseError("invalid offsets for observation replication", nil)
			}
			states[offset-1] = stateUnknown
			copy(states[offset:offset+splitIndex], states[:splitIndex])
			// copy brokenParts for the repeated observation
			start := numberOfBrokenGroups * (idx + 1)
			end := numberOfBrokenGroups * (idx + 2)
			if end > len(brokenParts) {
				return nil, aoc_utils.NewParseError("internal error copying broken parts", nil)
			}
			copy(brokenParts[start:end], brokenParts[:numberOfBrokenGroups])
		}
	}
	// sanity check that the sum of broken parts do not exceed the number of states
	totalBroken := 0
	for _, b := range brokenParts {
		totalBroken += b
	}
	if len(states) == 0 {
		return nil, aoc_utils.NewParseError("Empty line", nil)
	}
	if len(brokenParts) == 0 {
		return nil, aoc_utils.NewParseError("No broken parts", nil)
	}
	if totalBroken > len(states) {
		return nil, aoc_utils.NewParseError("Sum of broken parts exceeds line length", nil)
	}
	return &line{BrokenParts: brokenParts, States: states}, nil
}

func parseInput(r io.Reader, observationMultiplier int) ([]*line, error) {
	scanner := bufio.NewScanner(r)
	lines := make([]*line, 0, 32)
	for scanner.Scan() {
		l, err := parseLine(scanner.Bytes(), observationMultiplier)
		if err != nil {
			return nil, err
		}
		lines = append(lines, l)
	}
	return lines, nil
}
