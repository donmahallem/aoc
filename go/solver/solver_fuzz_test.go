package solver

import (
	_ "embed"
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/test_utils"
	"github.com/donmahallem/aoc/test/testgen/model"
)

// addTestDataSeeds adds every input from test/data.json as a seed for each part it covers.
func addTestDataSeeds(f *testing.F) {
	test_utils.EachTestCase(test_utils.LoadDefaultTestData(f), func(year, day int, name string, tc model.TestCase) {
		input, err := tc.Fetch()
		if err != nil {
			f.Logf("Skipping test case %d/%d (%s): %v", year, day, name, err)
			return
		}
		if tc.Part1 != nil {
			f.Add(uint8(year), uint8(day), false, string(input))
		}
		if tc.Part2 != nil {
			f.Add(uint8(year), uint8(day), true, string(input))
		}
	})
}

func FuzzSolver(f *testing.F) {
	// dynamically add tests from samples and full test files
	addTestDataSeeds(f)

	f.Fuzz(func(t *testing.T, year, day uint8, partTwo bool, input string) {

		reader := strings.NewReader(input)
		s := NewSolver()

		part := 1
		if partTwo {
			part = 2
		}
		// only run registered parts to avoid noise from unregistered selections
		if _, ok := s.GetRegistry().GetPart(aoc_utils.PartSelector{Year: int(year), Day: int(day), Part: part}); !ok {
			// count as skipped because the part is not registered
			t.Skip("Part not registered")
			return
		}
		res := s.Solve(int(year), int(day), part, reader)
		if res.Error != nil {
			if parseError, ok := errors.AsType[*aoc_utils.ParseError](res.Error); ok {
				t.Skipf("Parse error: %v", parseError)
				return
			} else if notImplError, ok := errors.AsType[*aoc_utils.NotImplementedError](res.Error); ok {
				t.Skipf("Not implemented: %v", notImplError)
				return
			} else if solverError, ok := errors.AsType[*aoc_utils.SolverError](res.Error); ok {
				t.Skipf("Solver error: %v", solverError)
				return
			} else if timeoutError, ok := errors.AsType[*aoc_utils.TimeoutError](res.Error); ok {
				t.Fatalf("Timeout error: %v", timeoutError)
			} else if numError, ok := errors.AsType[*strconv.NumError](res.Error); ok {
				t.Skipf("Atoi error: failed to parse \"%s\" to %s", numError.Num, numError.Func)
				return
			} else {
				t.Fatalf("Unexpected error: %v", res.Error)
			}
		}

	})

}
