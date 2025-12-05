package day05

import "sort"

func compressValidRanges(ranges *[]validRange) {
	rs := *ranges
	sort.Slice(rs, func(i, j int) bool { return rs[i].Min < rs[j].Min })

	i := 0
	for i < len(rs)-1 {
		currentRange := &rs[i]
		nextRange := &rs[i+1]
		if currentRange.Overlaps(*nextRange) || currentRange.Max+1 == nextRange.Min {
			if nextRange.Max > currentRange.Max {
				currentRange.Max = nextRange.Max
			}
			rs = append(rs[:i+1], rs[i+2:]...)
		} else {
			i++
		}
	}
	*ranges = rs
}
