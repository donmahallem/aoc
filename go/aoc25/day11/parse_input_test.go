package day11

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	reader := strings.NewReader(testData)
	res, err := parseInput(reader)
	if err != nil {
		t.Fatalf("parseInput returned an error: %v", err)
	}

	if len(res) != 10 {
		t.Errorf("unexpected number of nodes: got %d, want 10", len(res))
	}

	testCases := []struct {
		name            string
		key             uint64
		expectedTargets int
	}{
		{name: "you", key: uint64('y')<<16 + uint64('o')<<8 + uint64('u'), expectedTargets: 2},
		{name: "aaa", key: uint64('a')<<16 + uint64('a')<<8 + uint64('a'), expectedTargets: 2},
		{name: "ddd", key: uint64('d')<<16 + uint64('d')<<8 + uint64('d'), expectedTargets: 1},
		{name: "iii", key: uint64('i')<<16 + uint64('i')<<8 + uint64('i'), expectedTargets: 1},
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
