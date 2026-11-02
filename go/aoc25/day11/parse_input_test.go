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

	if res.ids.size() != 4 {
		t.Errorf("unexpected number of nodes: got %d, want 4", res.ids.size())
	}

	testCases := []struct {
		name            string
		key             uint32
		expectedTargets int
	}{
		{name: "aaa", key: encodeKey('a', 'a', 'a'), expectedTargets: 2},
		{name: "bbb", key: encodeKey('b', 'b', 'b'), expectedTargets: 1},
		{name: "ccc", key: encodeKey('c', 'c', 'c'), expectedTargets: 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			id, ok := res.ids.find(testCase.key)
			if !ok {
				t.Errorf("expected to find key %d, but not found", testCase.key)
				return
			}
			targets := res.successors(id)
			if len(targets) != testCase.expectedTargets {
				t.Errorf("unexpected number of targets for key %d: got %d, want %d", testCase.key, len(targets), testCase.expectedTargets)
			}
		})
	}

}

func Test_parseInput_rejectsNonThreeLetterNames(t *testing.T) {
	for _, input := range []string{"aa: bbb\n", "aaaa: bbb\n", "aaa: bb\n", "aaa: bbbb\n"} {
		if _, err := parseInput(strings.NewReader(input)); err == nil {
			t.Errorf("parseInput(%q) returned no error", input)
		}
	}
}
