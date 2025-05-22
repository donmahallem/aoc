package day05

import (
	"io"
	"math"
	"slices"
)

func Part2(in io.Reader) int {
	almanac := ParseAlmanac(in)
	return FindSmallestEndIndexFixed(almanac)
}

func IntersectInterval(a, b Interval) *Interval {
	start := max(b.Start, a.Start)
	end := a.End
	if b.End < end {
		end = b.End
	}
	if start < end {
		return &Interval{Start: start, End: end}
	}
	return nil
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.Start - b.Start
	})
	merged := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		if intervals[i].Start <= last.End { // overlapping or contiguous
			if intervals[i].End > last.End {
				last.End = intervals[i].End
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

func subtractUnion(a Interval, subs []Interval) []Interval {
	if len(subs) == 0 {
		return []Interval{a}
	}
	var result []Interval
	last := a.Start
	for _, sub := range subs {
		if sub.Start > last {
			result = append(result, Interval{Start: last, End: sub.Start})
		}
		if sub.End > last {
			last = sub.End
		}
	}
	if last < a.End {
		result = append(result, Interval{Start: last, End: a.End})
	}
	return result
}

func applyMapping(intervals []Interval, mappingRules []AlmanacRange) []Interval {
	var result []Interval
	for _, interval := range intervals {
		var transformed []Interval
		var covered []Interval
		for _, mappingRule := range mappingRules {
			if inter := IntersectInterval(interval, mappingRule.From); inter != nil {
				offset := mappingRule.To.Start - mappingRule.From.Start
				t := inter.Shift(offset)
				transformed = append(transformed, t)
				covered = append(covered, *inter)
			}
		}
		mergedCovered := mergeIntervals(covered)
		identityParts := subtractUnion(interval, mergedCovered)
		result = append(result, append(transformed, identityParts...)...)
	}
	return mergeIntervals(result)
}

func findFinalIntervals(alma Almanac) []Interval {
	var seedIntervals []Interval
	for i := 0; i < len(alma.Seeds); i += 2 {
		start := alma.Seeds[i]
		length := alma.Seeds[i+1]
		seedIntervals = append(seedIntervals, Interval{Start: start, End: start + length})
	}
	current := mergeIntervals(seedIntervals)
	stages := [][]AlmanacRange{
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

func FindSmallestEndIndexFixed(alma Almanac) int {
	finalIntervals := findFinalIntervals(alma)
	minFinal := math.MaxInt
	for _, iv := range finalIntervals {
		if iv.Start < minFinal {
			minFinal = iv.Start
		}
	}

	return minFinal
}
