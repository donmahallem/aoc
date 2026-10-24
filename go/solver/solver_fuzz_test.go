package solver

import (
	_ "embed"
	"errors"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/test_utils"
)

// addTestDataSeeds adds every input from test/data.json as a seed for each part it covers.
func addTestDataSeeds(f *testing.F) {
	data, err := test_utils.LoadTestData(test_utils.DefaultTestDataPath())
	if err != nil {
		f.Fatal(err)
	}
	for _, yearKey := range slices.Sorted(maps.Keys(data)) {
		year, err := strconv.ParseUint(yearKey, 10, 8)
		if err != nil {
			f.Fatal(err)
		}
		days := data[yearKey]
		for _, dayKey := range slices.Sorted(maps.Keys(days)) {
			day, err := strconv.ParseUint(dayKey, 10, 8)
			if err != nil {
				f.Fatal(err)
			}
			for _, tc := range days[dayKey] {
				input, found := loadInput(tc)
				if !found {
					// don't throw exception just skip this test case
					f.Logf("Skipping test case %s/%s (%s) because input not found", yearKey, dayKey, tc.Name)
					continue
				}
				if tc.Part1 != nil {
					f.Add(uint8(year), uint8(day), false, input)
				}
				if tc.Part2 != nil {
					f.Add(uint8(year), uint8(day), true, input)
				}
			}
		}
	}
}

// loadInput loads the input from either inline or file if present.
func loadInput(tc test_utils.TestCase) (string, bool) {
	if tc.Input != nil {
		return *tc.Input, true
	}
	sampleFilePath := test_utils.GetTestSampleFilePath(*tc.File)
	if _, err := os.Stat(sampleFilePath); !os.IsNotExist(err) {
		raw, err := os.ReadFile(sampleFilePath)
		return string(raw), err == nil
	}
	// if no inline data or file is available pass
	return "", false
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
