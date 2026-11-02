package day19

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	aocutil "github.com/donmahallem/aoc/go/aoc_utils/math"
)

type IntInterval = aocutil.Interval[int]

type ratingRange struct {
	X IntInterval
	M IntInterval
	A IntInterval
	S IntInterval
}

// returns the number of combinations in the rating range
func (rr ratingRange) combinations() int {
	sizes := []int{rr.X.Size(), rr.M.Size(), rr.A.Size(), rr.S.Size()}
	total := int(1)
	for _, s := range sizes {
		if s == 0 {
			return 0
		}
		total *= s
	}
	return total
}

// gets the interval for the given letter
func (rr ratingRange) get(letter byte) (IntInterval, bool) {
	switch letter {
	case 'x':
		return rr.X, true
	case 'm':
		return rr.M, true
	case 'a':
		return rr.A, true
	case 's':
		return rr.S, true
	default:
		return IntInterval{}, false
	}
}

// sets the interval for the given letter
func (rr *ratingRange) set(letter byte, iv IntInterval) bool {
	switch letter {
	case 'x':
		rr.X = iv
	case 'm':
		rr.M = iv
	case 'a':
		rr.A = iv
	case 's':
		rr.S = iv
	default:
		return false
	}
	return true
}

func splitGreater(r ratingRange, letter byte, threshold int) (ratingRange, bool, ratingRange, bool) {
	iv, ok := r.get(letter)
	if !ok || iv.Max <= threshold {
		return ratingRange{}, false, r, true
	}

	match := r
	matchInterval := IntInterval{
		Min: max(iv.Min, threshold+1),
		Max: iv.Max,
	}
	match.set(letter, matchInterval)

	if iv.Min > threshold {
		return match, true, ratingRange{}, false
	}

	rest := r
	restInterval := IntInterval{
		Min: iv.Min,
		Max: min(iv.Max, threshold),
	}
	rest.set(letter, restInterval)

	return match, true, rest, true
}

// splits the rating range into two ranges that match and don't match the less than condition
func splitLess(r ratingRange, letter byte, threshold int) (ratingRange, bool, ratingRange, bool) {
	currentInterval, ok := r.get(letter)
	if !ok || currentInterval.Min >= threshold {
		return ratingRange{}, false, r, true
	}

	match := r
	matchInterval := IntInterval{
		Min: currentInterval.Min,
		Max: min(currentInterval.Max, threshold-1),
	}
	match.set(letter, matchInterval)

	if currentInterval.Max < threshold {
		return match, true, ratingRange{}, false
	}

	rest := r
	restInterval := IntInterval{
		Min: max(currentInterval.Min, threshold),
		Max: currentInterval.Max,
	}
	rest.set(letter, restInterval)

	return match, true, rest, true
}

// recursively counts the number of accepted ranges.
// depth bounds the recursion to prevent infinite loops in cyclic workflows
func countAcceptedRanges(workflows map[string]workflow, current string, rng ratingRange, depth int) (int, error) {
	combinations := rng.combinations()
	if combinations == 0 {
		return 0, nil
	}

	switch current {
	case actionReject:
		return 0, nil
	case actionAccept:
		return combinations, nil
	}

	if depth >= len(workflows) {
		return 0, aoc_utils.NewSolverErrorf("workflow %q is part of a cycle", current)
	}

	wf, ok := workflows[current]
	if !ok {
		return 0, nil
	}

	remaining := rng
	total := 0

	for _, rule := range wf.rules {
		var target string
		var match ratingRange
		var hasMatch, hasRest bool
		var rest ratingRange

		switch rule := rule.(type) {
		case *workflowRuleGreater:
			match, hasMatch, rest, hasRest = splitGreater(remaining, rule.letter, rule.value)
			target = rule.target
		case *workflowRuleLess:
			match, hasMatch, rest, hasRest = splitLess(remaining, rule.letter, rule.value)
			target = rule.target
		case *workflowRuleDirect:
			n, err := countAcceptedRanges(workflows, rule.target, remaining, depth+1)
			if err != nil {
				return 0, err
			}
			return total + n, nil
		}

		if hasMatch {
			n, err := countAcceptedRanges(workflows, target, match, depth+1)
			if err != nil {
				return 0, err
			}
			total += n
		}
		if !hasRest {
			return total, nil
		}
		remaining = rest
	}

	return total, nil
}

func Part2(r io.Reader) (int, error) {
	data := parseInput(r)

	initial := ratingRange{
		X: IntInterval{Min: 1, Max: 4000},
		M: IntInterval{Min: 1, Max: 4000},
		A: IntInterval{Min: 1, Max: 4000},
		S: IntInterval{Min: 1, Max: 4000},
	}

	return countAcceptedRanges(data.Workflows, "in", initial, 0)
}
