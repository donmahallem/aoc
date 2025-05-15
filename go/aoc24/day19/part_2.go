package day19

import (
	"io"
)

type ScoreCache = map[string]uint

func BisectSearchPatterns(patterns TowelPatterns, scoreCache *ScoreCache, subPattern Towel) uint {
	n := len(subPattern)
	if n == 1 {
		if _, ok := patterns[subPattern]; ok {
			return 1
		}
		return 0
	} else if n == 0 {
		return 1
	}
	if val, exists := (*scoreCache)[subPattern]; exists {
		return val
	}

	midpoint := n / 2
	total := BisectSearchPatterns(patterns, scoreCache, subPattern[:midpoint]) * BisectSearchPatterns(patterns, scoreCache, subPattern[midpoint:])

	for pattern := range patterns {
		towelLen := len(pattern)
		if towelLen > 1 && towelLen <= n {
			for i := range towelLen - 1 {
				start := midpoint - 1 - i
				end := start + towelLen
				if start < 0 || end > n {
					continue
				}
				if subPattern[start:end] == pattern {
					total += BisectSearchPatterns(patterns, scoreCache, subPattern[:start]) * BisectSearchPatterns(patterns, scoreCache, subPattern[end:])
				}
			}
		}
	}
	(*scoreCache)[subPattern] = total
	return total
}
func Part2(in io.Reader) uint {
	patterns, towls, _ := ParseInput(in)
	var count uint = 0
	cache := make(ScoreCache)
	for _, towl := range towls {
		count += BisectSearchPatterns(patterns, &cache, towl)

	}
	return count
}
