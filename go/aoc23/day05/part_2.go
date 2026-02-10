package day05

import (
	"io"
	"math"
	"slices"
)

func Part2(in io.Reader) (int, error) {
	almanac := parseAlmanac(in)
	return FindSmallestEndIndexFixed(almanac), nil
}

func mergeIntervals(intervals []almanacInterval) []almanacInterval {
	if len(intervals) == 0 {
		return intervals
	}
	slices.SortFunc(intervals, func(a, b almanacInterval) int {
		return a.Min - b.Min
	})
	merged := []almanacInterval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		if intervals[i].Min <= last.Max { // overlapping or contiguous
			if intervals[i].Max > last.Max {
				last.Max = intervals[i].Max
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

func subtractUnion(a almanacInterval, subs []almanacInterval) []almanacInterval {
	if len(subs) == 0 {
		return []almanacInterval{a}
	}
	var result []almanacInterval
	last := a.Min
	for _, sub := range subs {
		if sub.Min > last {
			result = append(result, almanacInterval{Min: last, Max: sub.Min})
		}
		if sub.Max > last {
			last = sub.Max
		}
	}
	if last < a.Max {
		result = append(result, almanacInterval{Min: last, Max: a.Max})
	}
	return result
}

func applyMapping(intervals []almanacInterval, mappingRules []almanacRange) []almanacInterval {
	var result []almanacInterval
	for _, interval := range intervals {
		var transformed []almanacInterval
		var covered []almanacInterval
		for _, mappingRule := range mappingRules {
			if inter, ok := interval.Intersection(mappingRule.From); ok {
				offset := mappingRule.To.Min - mappingRule.From.Min
				t := inter.Shift(offset)
				transformed = append(transformed, t)
				covered = append(covered, inter)
			}
		}
		mergedCovered := mergeIntervals(covered)
		identityParts := subtractUnion(interval, mergedCovered)
		result = append(result, append(transformed, identityParts...)...)
	}
	return mergeIntervals(result)
}

func findFinalIntervals(alma almanac) []almanacInterval {
	var seedIntervals []almanacInterval
	for i := 0; i < len(alma.Seeds); i += 2 {
		start := alma.Seeds[i]
		length := alma.Seeds[i+1]
		seedIntervals = append(seedIntervals, almanacInterval{Min: start, Max: start + length})
	}
	current := mergeIntervals(seedIntervals)
	stages := [][]almanacRange{
		alma.SeedToSoil,
		alma.SoilToFertilizer,
		alma.FertilizerToWater,
		alma.WaterToLight,
		alma.LightToTemperature,
		alma.TemperatureToHumidity,
		alma.HumidityToLocation,
	}
	for _, stage := range stages {
		current = applyMapping(current, stage)
	}
	return current
}

func FindSmallestEndIndexFixed(alma almanac) int {
	finalIntervals := findFinalIntervals(alma)
	minFinal := math.MaxInt
	for _, iv := range finalIntervals {
		if iv.Min < minFinal {
			minFinal = iv.Min
		}
	}

	return minFinal
}
