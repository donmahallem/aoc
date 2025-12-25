package solver

import (
	_ "embed"
	"errors"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

//go:embed sample1.txt
var sample_1 string

//go:embed sample2.txt
var sample_2 string

func FuzzSolver(f *testing.F) {
	f.Add(uint8(24), uint8(23), false, sample_1)
	f.Add(uint8(23), uint8(23), true, sample_2)
	f.Fuzz(func(t *testing.T, year, day uint8, partTwo bool, input string) {
		reader := strings.NewReader(input)
		s := NewSolver()

		part := 1
		if partTwo {
			part = 2
		}
		// only run registered parts to avoid noise from unregistered selections
		if _, ok := s.GetRegistry().GetPart(aoc_utils.PartSelector{Year: int(year), Day: int(day), Part: part}); !ok {
			return
		}
		res := s.Solve(int(year), int(day), part, reader)
		if res.Error != nil {
			var parseError *aoc_utils.ParseError
			var notImplError *aoc_utils.NotImplementedError
			var solverError *aoc_utils.SolverError
			var timeoutError *aoc_utils.TimeoutError
			if errors.As(res.Error, &parseError) {
				t.Logf("Parse error: %v", parseError)
				return
			} else if errors.As(res.Error, &notImplError) {
				t.Logf("Not implemented: %v", res.Error)
				return
			} else if errors.As(res.Error, &solverError) {
				t.Logf("Solver error: %v", res.Error)
				return
			} else if errors.As(res.Error, &timeoutError) {
				t.Fatalf("Timeout error: %v", res.Error)
			} else {
				t.Fatalf("Unexpected error: %v", res.Error)
			}
		}

	})

}
