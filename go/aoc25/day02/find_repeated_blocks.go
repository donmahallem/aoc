package day02

import "github.com/donmahallem/aoc/go/aoc_utils/int_util"

// Assuming intInterval is defined elsewhere like this:
// type intInterval struct {
// 	Min int
// 	Max int
// }

func findRepeatedBlocks(val intInterval) map[uint64]uint64 {
	result := make(map[uint64]uint64)
	lbound := val.Min
	ubound := val.Max

	minLen := int_util.Log10Int(lbound) + 1
	maxLen := int_util.Log10Int(ubound) + 1

	for totalLength := minLen; totalLength <= maxLen; totalLength++ {
		for blockLen := uint64(1); blockLen <= totalLength/2; blockLen++ {
			if totalLength%blockLen != 0 {
				continue
			}

			k := totalLength / blockLen
			if k < 2 {
				continue
			}

			minBlock := int_util.IntPow(10, blockLen-1)
			maxBlock := int_util.IntPow(10, blockLen) - 1

			// calc  multiplier
			numerator := int_util.IntPow(10, blockLen*k) - 1
			denominator := int_util.IntPow(10, blockLen) - 1
			mult := numerator / denominator

			for block := minBlock; block <= maxBlock; block++ {
				v := block * mult

				if v < lbound {
					continue
				}
				if v > ubound {
					// as block increases, v will only get larger
					break
				}

				// Ensure v has exactly total_length digits
				lowerLimit := int_util.IntPow(10, totalLength-1)
				upperLimit := int_util.IntPow(10, totalLength)

				if v >= lowerLimit && v < upperLimit {
					result[uint64(v)] = uint64(k)
				}
			}
		}
	}

	return result
}
