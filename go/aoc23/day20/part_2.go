package day20

import (
	"io"

	"github.com/donmahallem/aoc/aoc_utils/math/lcm"
)

const rxModuleId int = 'r'<<8 + 'x'

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

func HandlePart2(in io.Reader, outputId int) int {

	inp := ParseInput(in)

	// getting previous conjuction module
	rxSources := findModulesTargetingRx(&inp, outputId)

	if len(rxSources) == 0 {
		panic("no modules target rx")
	}
	targetID := rxSources[0]
	conj, ok := inp.Modules[targetID].(*conjunctionModule)
	if !ok || len(conj.inputs) == 0 {
		panic("rx source is not a conjunction module or has no inputs")
	}
	periods := make(map[int]int, len(conj.inputs))
	needed := len(conj.inputs)

	for press := 1; len(periods) < needed; press++ {
		currentPress := press
		simulatePress(&inp, func(p pulse) {
			if p.to == targetID && p.high {
				if _, recorded := periods[p.from]; !recorded {
					periods[p.from] = currentPress
				}
			}
		})
	}
	var result int = 1
	for _, period := range periods {
		result = lcm.LcmInt(result, period)
	}

	return result
}

func Part2(in io.Reader) int {
	return HandlePart2(in, rxModuleId)
}
