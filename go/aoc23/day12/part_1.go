package day12

import (
	"io"
)

func solveLine(l line) (int, error) {
	states := l.States
	hints := l.BrokenParts

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

	return countPossiblePermutations(states, hints, remainingBroken, 0, 0, 0, memo, visited, strideHint, strideRun, stateLen, hintLen), nil
}

func Part1(in io.Reader) (int, error) {
	lines, err := parseInput(in, 1)
	if err != nil {
		return 0, err
	}
	var total int = 0
	for _, l := range lines {
		stepCount, err := solveLine(*l)
		if err != nil {
			return 0, err
		}
		total += stepCount
	}
	return total, nil
}

func countPossiblePermutations(states []springState, hints []int, remaining []int, idx, hintIdx, run int,
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
	case stateBroken:
		if hintIdx < hintLen && run < hints[hintIdx] {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx, run+1, memo, visited, strideHint, strideRun, stateLen, hintLen)
		}
	case stateOk:
		if run == 0 {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
		} else if hints[hintIdx] == run {
			total = countPossiblePermutations(states, hints, remaining, idx+1, hintIdx+1, 0, memo, visited, strideHint, strideRun, stateLen, hintLen)
		}
	case stateUnknown:
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
