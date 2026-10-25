package day20

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

const rxModuleId int = 'r'<<8 + 'x'

// bounds the press search
const maxPresses = 1 << 16

// search for modules targeting the rx module
func findModulesTargetingRx(inp *parsedInput, outputId int) []int {
	ids := make([]int, 0, 4)
	for id, mod := range inp.Modules {
		for _, target := range mod.TargetIds() {
			if target == outputId {
				ids = append(ids, id)
				break
			}
		}
	}
	return ids
}

func HandlePart2(in io.Reader, outputId int) (int, error) {

	inp := ParseInput(in)

	// getting previous conjuction module
	rxSources := findModulesTargetingRx(&inp, outputId)

	if len(rxSources) == 0 {
		// malformed or incomplete circuit: treat as non-fatal and return 0
		return 0, nil
	}
	targetID := rxSources[0]
	conj, ok := inp.Modules[targetID].(*conjunctionModule)
	if !ok || len(conj.inputs) == 0 {
		// similarly, return 0 for incomplete/invalid configurations
		return 0, nil
	}
	periods := make(map[int]int, len(conj.inputs))
	needed := len(conj.inputs)

	// some circuits never send every input a high pulse
	for press := 1; len(periods) < needed; press++ {
		if press > maxPresses {
			return 0, aoc_utils.NewSolverErrorf("no cycle found within %d presses", maxPresses)
		}
		currentPress := press
		_, _, err := simulatePress(&inp, func(p pulse) {
			if p.to == targetID && p.high {
				if _, recorded := periods[p.from]; !recorded {
					periods[p.from] = currentPress
				}
			}
		})
		if err != nil {
			return 0, err
		}
	}
	var result int = 1
	for _, period := range periods {
		result = int_util.LcmInt(result, period)
	}

	return result, nil
}

func Part2(in io.Reader) (int, error) {
	return HandlePart2(in, rxModuleId)
}
