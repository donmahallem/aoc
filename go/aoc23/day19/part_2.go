package day19

import (
	"io"

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

// recursively counts the number of accepted ranges
func countAcceptedRanges(workflows map[string]workflow, current string, rng ratingRange) int {
	combinations := rng.combinations()
	if combinations == 0 {
		return 0
	}

	switch current {
	case actionReject:
		return 0
	case actionAccept:
		return combinations
	}

	wf, ok := workflows[current]
	if !ok {
		return 0
	}

	remaining := rng
	total := 0

	for _, rule := range wf.rules {
		switch rule := rule.(type) {
		case *workflowRuleGreater:
			match, hasMatch, rest, hasRest := splitGreater(remaining, rule.letter, rule.value)
			if hasMatch {
				total += countAcceptedRanges(workflows, rule.target, match)
			}
			if !hasRest {
				return total
			}
			remaining = rest
		case *workflowRuleLess:
			match, hasMatch, rest, hasRest := splitLess(remaining, rule.letter, rule.value)
			if hasMatch {
				total += countAcceptedRanges(workflows, rule.target, match)
			}
			if !hasRest {
				return total
			}
			remaining = rest
		case *workflowRuleDirect:
			total += countAcceptedRanges(workflows, rule.target, remaining)
			return total
		}
	}

	return total
}

func Part2(in io.Reader) int {
	data := ParseInput(in)

	initial := ratingRange{
		X: IntInterval{Min: 1, Max: 4000},
		M: IntInterval{Min: 1, Max: 4000},
		A: IntInterval{Min: 1, Max: 4000},
		S: IntInterval{Min: 1, Max: 4000},
	}

	return countAcceptedRanges(data.Workflows, "in", initial)
}
