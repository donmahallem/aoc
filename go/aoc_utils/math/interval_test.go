package math_test

import (
	"strconv"
	"testing"

	"github.com/donmahallem/aoc/aoc_utils/math"
)

func TestIntervalSize(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		interval math.Interval[int]
		expected int
	}{
		"positive range": {interval: math.Interval[int]{Min: 1, Max: 5}, expected: 5},
		"single point":   {interval: math.Interval[int]{Min: 3, Max: 3}, expected: 1},
		"empty":          {interval: math.Interval[int]{Min: 5, Max: 3}, expected: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tc.interval.Size(); got != tc.expected {
				t.Fatalf("Size() = %d, expected %d", got, tc.expected)
			}
		})
	}
}

func TestIntervalContains(t *testing.T) {
	t.Parallel()

	iv := math.Interval[int]{Min: 2, Max: 6}

	cases := map[int]bool{
		1: false,
		2: true,
		4: true,
		6: true,
		7: false,
	}

	for value, expected := range cases {
		value, expected := value, expected
		t.Run("value_"+strconv.Itoa(value), func(t *testing.T) {
			if got := iv.Contains(value); got != expected {
				t.Fatalf("Contains(%d) = %v, expected %v", value, got, expected)
			}
		})
	}
}

func TestIntervalOverlaps(t *testing.T) {
	t.Parallel()

	iv := math.Interval[int]{Min: 5, Max: 10}

	tests := map[string]struct {
		other    math.Interval[int]
		expected bool
	}{
		"overlap middle": {other: math.Interval[int]{Min: 8, Max: 12}, expected: true},
		"touch lower":    {other: math.Interval[int]{Min: 1, Max: 5}, expected: true},
		"touch upper":    {other: math.Interval[int]{Min: 10, Max: 15}, expected: true},
		"disjoint lower": {other: math.Interval[int]{Min: 0, Max: 4}, expected: false},
		"disjoint upper": {other: math.Interval[int]{Min: 11, Max: 20}, expected: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := iv.Overlaps(tc.other); got != tc.expected {
				t.Fatalf("Overlaps(%+v) = %v, expected %v", tc.other, got, tc.expected)
			}
		})
	}
}

func TestIntervalIntersection(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		first         math.Interval[int]
		second        math.Interval[int]
		expectedIV    math.Interval[int]
		expectedFound bool
	}{
		"regular overlap": {
			first:         math.Interval[int]{Min: 1, Max: 5},
			second:        math.Interval[int]{Min: 3, Max: 7},
			expectedIV:    math.Interval[int]{Min: 3, Max: 5},
			expectedFound: true,
		},
		"touching": {
			first:         math.Interval[int]{Min: 0, Max: 2},
			second:        math.Interval[int]{Min: 2, Max: 4},
			expectedIV:    math.Interval[int]{Min: 2, Max: 2},
			expectedFound: true,
		},
		"no intersection": {
			first:         math.Interval[int]{Min: 8, Max: 9},
			second:        math.Interval[int]{Min: 10, Max: 12},
			expectedFound: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotIV, ok := tc.first.Intersection(tc.second)
			if ok != tc.expectedFound {
				t.Fatalf("Intersection found=%v, expected %v", ok, tc.expectedFound)
			}
			if ok && gotIV != tc.expectedIV {
				t.Fatalf("Intersection = %+v, expected %+v", gotIV, tc.expectedIV)
			}
		})
	}
}

func TestIntervalFix(t *testing.T) {
	t.Parallel()

	iv := math.Interval[int]{Min: 9, Max: 4}
	iv.Fix()

	if iv.Min != 4 || iv.Max != 9 {
		t.Fatalf("Fix() => Interval{Min:%d, Max:%d}, expected Min:4 Max:9", iv.Min, iv.Max)
	}
}
