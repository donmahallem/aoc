package day11

import (
	"strings"
	"testing"
)

// Minimal inline sample graph
const testData = "aaa: bbb ccc\nbbb: out\nccc: out\n"

func Test_parseInput(t *testing.T) {
	reader := strings.NewReader(testData)
	res, err := parseInput(reader)
	if err != nil {
		t.Fatalf("parseInput returned an error: %v", err)
	}

	if len(res) != 3 {
		t.Errorf("unexpected number of nodes: got %d, want 3", len(res))
	}

	testCases := []struct {
		name            string
		key             uint64
		expectedTargets int
	}{
		{name: "aaa", key: uint64('a')<<16 + uint64('a')<<8 + uint64('a'), expectedTargets: 2},
		{name: "bbb", key: uint64('b')<<16 + uint64('b')<<8 + uint64('b'), expectedTargets: 1},
		{name: "ccc", key: uint64('c')<<16 + uint64('c')<<8 + uint64('c'), expectedTargets: 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			targets, ok := res[testCase.key]
			if !ok {
				t.Errorf("expected to find key %d, but not found", testCase.key)
				return
			}
			if len(targets) != testCase.expectedTargets {
				t.Errorf("unexpected number of targets for key %d: got %d, want %d", testCase.key, len(targets), testCase.expectedTargets)
			}
		})
	}

}
