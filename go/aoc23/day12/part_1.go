package day12

import (
	"bufio"
	"bytes"
	"io"

	bytesutil "github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

type SpringState uint8

const (
	StateUnknown SpringState = '?'
	StateOk      SpringState = '.'
	StateBroken  SpringState = '#'
)

type Line struct {
	BrokenParts []int
	States      []SpringState
}

/**
 * Parses a single line of input into a Line struct.
 * observationMultiplier determines how many times the observation is repeated.
 */
func parseLine(line []byte, observationMultiplier int) Line {
	splitIndex := bytes.IndexByte(line, ' ')
	brokenPartSeries := bytes.Split(line[splitIndex+1:], []byte{','})
	numberOfBrokenGroups := len(brokenPartSeries)
	brokenParts := make([]int, numberOfBrokenGroups*observationMultiplier)
	for i := range numberOfBrokenGroups {
		brokenParts[i] = bytesutil.ByteSequenceToInt[int](brokenPartSeries[i])
	}
	states := make([]SpringState, (splitIndex*observationMultiplier)+(observationMultiplier-1))
	for idx, c := range line[:splitIndex] {
		switch c {
		case '?':
			states[idx] = StateUnknown
		case '#':
			states[idx] = StateBroken
		case '.':
			states[idx] = StateOk
		}
	}
	if observationMultiplier > 1 {
		for idx := range observationMultiplier - 1 {
			offset := (splitIndex + 1) * (idx + 1)
			states[offset-1] = StateUnknown
			copy(states[offset:offset+splitIndex], states[:splitIndex])
			copy(brokenParts[numberOfBrokenGroups*(idx+1):numberOfBrokenGroups*(idx+2)], brokenParts[:numberOfBrokenGroups])
		}
	}
	return Line{BrokenParts: brokenParts, States: states}
}

func ParseInput(r io.Reader, observationMultiplier int) []Line {
	scanner := bufio.NewScanner(r)
	lines := make([]Line, 0, 32)
	for scanner.Scan() {
		lines = append(lines, parseLine(scanner.Bytes(), observationMultiplier))
	}
	return lines
}

func SolveLine(line Line) int {
	states := line.States
	hints := line.BrokenParts

	stateLen := len(states)
	hintLen := len(hints)

	var maxHint int
	remainingBroken := make([]int, hintLen+1)
	for i := hintLen - 1; i >= 0; i-- {
		if hints[i] > maxHint {
			maxHint = hints[i]
		}
		remainingBroken[i] = remainingBroken[i+1] + hints[i]
	}

	strideRun := maxHint + 1
	strideHint := (hintLen + 1) * strideRun
	memoSize := (stateLen + 1) * strideHint

	memo := make([]int, memoSize)
	visited := make([]bool, memoSize)

	return countPossiblePermutations(states, hints, remainingBroken, 0, 0, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
}

func Part1(in io.Reader) int {
	lines := ParseInput(in, 1)
	var total int = 0
	for _, line := range lines {
		total += SolveLine(line)
	}
	return total
}

func countPossiblePermutations(states []SpringState, hints []int, remaining []int, idx, hintIdx, run int,
	memo []int, visited []bool, strideHint, strideRun, stateLen, hintLen int) int {
	key := idx*strideHint + hintIdx*strideRun + run
	if visited[key] {
		return memo[key]
	}
	visited[key] = true

	if remaining[hintIdx]-run > stateLen-idx {
		memo[key] = 0
		return 0
	}

	if idx == stateLen {
		if run > 0 {
			if hintIdx >= hintLen || hints[hintIdx] != run {
				memo[key] = 0
				return 0
			}
			hintIdx++
		}
		if hintIdx == hintLen {
			memo[key] = 1
		} else {
			memo[key] = 0
		}
		return memo[key]
	}

	state := states[idx]
	if run > 0 {
		if hintIdx >= hintLen || run > hints[hintIdx] {
			memo[key] = 0
			return 0
		}
	}

	var total int

	switch state {
	case StateBroken:
		if hintIdx < hintLen && run < hints[hintIdx] {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx, run+1, memo, visited, strideHint, strideRun, stateLen, hintLen)
		}
	case StateOk:
		if run == 0 {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
		} else if hints[hintIdx] == run {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx+1, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
		}
	case StateUnknown:
		if hintIdx < hintLen && run < hints[hintIdx] {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx, run+1, memo, visited, strideHint, strideRun, stateLen, hintLen)
		}
		if run == 0 {
			total += countPossiblePermutations(states, hints, remaining, idx+1, hintIdx, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
		} else if hints[hintIdx] == run {
			total += countPossiblePermutations(states, hints, remaining, idx+1, hintIdx+1, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
		}
	}

	memo[key] = total
	return total
}
