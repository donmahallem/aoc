package day07

import (
	"testing"
)

func Test_parseLine(t *testing.T) {

	testCases := []struct {
		rawLine        []byte
		expectedResult int
		expectedTerms  []int
		expectError    bool
	}{
		{[]byte("190: 10 19"), 190, []int{10, 19}, false},
		{[]byte("42: 7 14 21"), 42, []int{7, 14, 21}, false},
		{[]byte("7: 3"), 7, []int{3}, false},
		{[]byte("1000: 100 200 300 400"), 1000, []int{100, 200, 300, 400}, false},
		{[]byte("1                "), 0, nil, true},
		{[]byte("123 456"), 0, nil, true},  // Missing colon
		{[]byte("abc: 123"), 0, nil, true}, // Invalid target
		{[]byte("123: abc"), 0, nil, true}, // Invalid value
	}
	for _, tc := range testCases {
		parsedLine := &parsedLineData{}
		err := parseLine(tc.rawLine, parsedLine)

		if tc.expectError {
			if err == nil {
				t.Errorf("Expected error for input '%s', but got nil", string(tc.rawLine))
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input '%s': %v", string(tc.rawLine), err)
			continue
		}

		if parsedLine.Result != tc.expectedResult {
			t.Errorf(`Expected result to be %d, got %d`, tc.expectedResult, parsedLine.Result)
		}
		if len(parsedLine.TestValues) != len(tc.expectedTerms) {
			t.Errorf(`Expected number of terms to be %d, got %d`, len(tc.expectedTerms), len(parsedLine.TestValues))
		}
		for i, term := range tc.expectedTerms {
			if parsedLine.TestValues[i] != term {
				t.Errorf(`Expected term %d to be %d, got %d`, i, term, parsedLine.TestValues[i])
			}
		}
	}
}
