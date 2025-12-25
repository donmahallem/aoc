package day05

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils/math"
)

type validRange = math.Interval[uint64]

type taskData struct {
	validRanges []validRange
	ingredients []uint64
}

func parseInput(in io.Reader) (taskData, error) {
	scanner := bufio.NewScanner(in)
	var data taskData
	inRangesBlock := true
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			inRangesBlock = false
			continue
		}
		if inRangesBlock {
			currentInterval := validRange{}
			currentNumber := &currentInterval.Min
			for i := range len(line) {
				b := line[i]
				if b >= '0' && b <= '9' {
					*currentNumber = *currentNumber*10 + uint64(b-'0')
				} else if b == '-' {
					currentNumber = &currentInterval.Max
				}
			}
			data.validRanges = append(data.validRanges, currentInterval)
		} else {
			currentNumber := uint64(0)
			for i := range len(line) {
				b := line[i]
				if b >= '0' && b <= '9' {
					currentNumber = currentNumber*10 + uint64(b-'0')
				}
			}
			data.ingredients = append(data.ingredients, currentNumber)
		}
	}
	return data, nil
}
