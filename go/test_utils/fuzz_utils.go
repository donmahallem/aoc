package test_utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func AddTestDataSeeds(f *testing.F, year, day int) {
	yearKey := fmt.Sprintf("%02d", year)
	dayKey := fmt.Sprintf("%02d", day)
	data, err := LoadTestData(DefaultTestDataPath())
	if err != nil {
		f.Fatal(err)
	}
	for _, tc := range data[yearKey][dayKey] {
		input, found := loadTestCaseInput(tc)
		if !found {
			f.Logf("Skipping test case %s/%s (%s) because input not found", yearKey, dayKey, tc.Name)
			continue
		}
		f.Add(input)
	}
}

func loadTestCaseInput(tc TestCase) (string, bool) {
	if tc.Input != nil {
		return *tc.Input, true
	}
	raw, err := os.ReadFile(GetTestSampleFilePath(*tc.File))
	return string(raw), err == nil
}

// FuzzHandleError skips expected solver errors and fails the fuzz run on anything else.
func FuzzHandleError(t *testing.T, err error) bool {
	t.Helper()
	if err == nil {
		return true
	}
	if parseError, ok := errors.AsType[*aoc_utils.ParseError](err); ok {
		t.Skipf("Parse error: %v", parseError)
		return true
	} else if notImplError, ok := errors.AsType[*aoc_utils.NotImplementedError](err); ok {
		t.Skipf("Not implemented: %v", notImplError)
		return true
	} else if solverError, ok := errors.AsType[*aoc_utils.SolverError](err); ok {
		t.Skipf("Solver error: %v", solverError)
		return true
	} else if timeoutError, ok := errors.AsType[*aoc_utils.TimeoutError](err); ok {
		t.Fatalf("Timeout error: %v", timeoutError)
		return false
	} else if numError, ok := errors.AsType[*strconv.NumError](err); ok {
		t.Skipf("Atoi error: failed to parse \"%s\" to %s", numError.Num, numError.Func)
		return true
	} else {
		t.Fatalf("Unexpected error: %v", err)
		return false
	}
}
