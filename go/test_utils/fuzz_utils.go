package test_utils

import (
	"bufio"
	"errors"
	"maps"
	"slices"
	"strconv"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func AddTestDataSeeds(f *testing.F, year, day int) {
	cases := LoadDefaultTestData(f)[year][day]
	for _, name := range slices.Sorted(maps.Keys(cases)) {
		tc := cases[name]
		input, err := tc.Fetch()
		if err != nil {
			f.Logf("Skipping test case %d/%d (%s): %v", year, day, name, err)
			continue
		}
		f.Add(string(input))
	}
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
	} else if errors.Is(err, bufio.ErrTooLong) {
		t.Skipf("Bufio error: %v", err)
		return true
	} else {
		t.Fatalf("Unexpected error: %v", err)
		return false
	}
}
