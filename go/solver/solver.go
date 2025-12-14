package solver

import (
	"fmt"
	"io"
	"reflect"
	"time"

	"github.com/donmahallem/aoc/go/aoc23"
	"github.com/donmahallem/aoc/go/aoc24"
	"github.com/donmahallem/aoc/go/aoc25"
	"github.com/donmahallem/aoc/go/aoc_utils"
)

type SolveResult struct {
	Result   string
	Duration time.Duration
	Error    error
}

type Solver struct {
	registry aoc_utils.Registry
}

func NewSolver() *Solver {
	partRegistry := aoc_utils.NewRegistry()
	aoc23.RegisterParts(&partRegistry)
	aoc24.RegisterParts(&partRegistry)
	aoc25.RegisterParts(&partRegistry)
	return &Solver{registry: partRegistry}
}

func (s *Solver) GetRegistry() *aoc_utils.Registry {
	return &s.registry
}

func (s *Solver) Solve(year, day, part int, input io.Reader) SolveResult {
	partSelector := aoc_utils.PartSelector{
		Year: year,
		Day:  day,
		Part: part,
	}

	takeFun, ok := s.registry.GetPart(partSelector)
	if !ok {
		return SolveResult{Error: fmt.Errorf("could not find requested part %v", partSelector)}
	}

	function := reflect.ValueOf(takeFun)
	startTime := time.Now()
	results := function.Call([]reflect.Value{reflect.ValueOf(input)})
	endTime := time.Now()

	res := results[0].Interface()
	var resStr string

	switch v := res.(type) {
	case int, uint, int32, uint32, uint16, uint8, uint64, int16, int8, int64:
		resStr = fmt.Sprintf("%d", v)
	case string:
		resStr = v
	case []int:
		for i, val := range v {
			if i > 0 {
				resStr += ","
			}
			resStr += fmt.Sprintf("%d", val)
		}
	case aoc_utils.Point[int16]:
		resStr = fmt.Sprintf("{X:%d,Y:%d}", v.X, v.Y)
	default:
		resStr = fmt.Sprintf("%v", v)
	}

	return SolveResult{
		Result:   resStr,
		Duration: endTime.Sub(startTime),
	}
}
