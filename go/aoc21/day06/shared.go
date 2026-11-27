package day06

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

const maxLineLength = 786

func caclcuateDivisions(counts [9]int, maxDepth int) int {
	currentIdx := 0
	for range maxDepth {
		spawningCount := counts[currentIdx]
		// current entries with value '0' are spawning new fish
		counts[currentIdx] = 0
		// shift from zero to pos 6. as we shift it needs to be 6+1
		counts[(currentIdx+7)%9] += spawningCount
		// shift from zero to pos 8. as we shift it next step it needs to be 8+1
		counts[(currentIdx+9)%9] += spawningCount
		currentIdx = (currentIdx + 1) % 9
	}

	totalSum := 0
	for _, v := range counts {
		totalSum += v
	}
	return totalSum
}

func parseInput(in io.Reader) ([9]int, error) {
	data := make([]uint8, maxLineLength)
	counts := [9]int{}
	ln, err := in.Read(data)
	if err != nil && err != io.EOF {
		return [9]int{}, err
	}
	for i := range data[:ln] {
		switch data[i] {
		case '\n', '\r':
			return counts, nil
		case ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8':
			counts[int(data[i]-'0')]++
		default:
			return [9]int{}, aoc_utils.NewUnexpectedInputError(data[i])
		}
	}
	return counts, nil
}
