package day09

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func convertInput(inp []byte) ([]int16, error) {
	baseData := make([]int16, 0)
	for i := range inp {
		if inp[i] < '0' || inp[i] > '9' {
			return nil, aoc_utils.NewUnexpectedInputError(inp[i])
		}
		spec := inp[i] - '0'
		for range spec {
			if i%2 == 0 {
				baseData = append(baseData, int16((i / 2)))
			} else {
				baseData = append(baseData, -1)
			}
		}
	}
	return baseData, nil
}
func checkSum(data *[]int16) int {
	checkSum := 0
	for i := range len(*data) {
		if (*data)[i] >= 0 {
			checkSum += int((*data)[i]) * i
		}
	}
	return checkSum
}
func compactData(data *[]int16) {
	j := len(*data) - 1
	for i := range len(*data) {
		if (*data)[i] >= 0 {
			continue
		} else if i >= j {
			break
		}
		for ; j > i; j-- {
			if (*data)[j] >= 0 {
				(*data)[i] = (*data)[j]
				(*data)[j] = -1
				break
			}
		}
	}
}

func Part1(in io.Reader) (int, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return 0, err
	}
	if len(data) == 0 {
		return 0, nil
	}
	// trim
	endIdx := len(data) - 1
	for endIdx >= 0 && (data[endIdx] == '\n' || data[endIdx] == '\r') {
		endIdx--
	}

	// endIdx must sit on a file entry (even index).
	if endIdx%2 != 0 {
		endIdx--
	}

	endFileID := endIdx / 2
	endRemaining := int(data[endIdx] - '0')

	positionidx := 0
	totalValue := 0

	for frontIdx := 0; frontIdx <= endIdx; frontIdx++ {
		if data[frontIdx] < '0' || data[frontIdx] > '9' {
			return 0, aoc_utils.NewUnexpectedInputError(data[frontIdx])
		}
		frontValue := int(data[frontIdx] - '0')

		if frontIdx%2 == 0 {
			// File section: place blocks that haven't been relocated.
			fileID := frontIdx / 2
			if frontIdx == endIdx {
				// Take care of remaining blocks until end
				for range endRemaining {
					totalValue += fileID * positionidx
					positionidx++
				}
				break
			}
			for range frontValue {
				totalValue += fileID * positionidx
				positionidx++
			}
		} else {
			// Free space: fill one block at a time from the end.
			for space := frontValue; space > 0; space-- {
				for endRemaining == 0 {
					endIdx -= 2
					if endIdx <= frontIdx {
						return totalValue, nil
					}
					if data[endIdx] < '0' || data[endIdx] > '9' {
						return 0, aoc_utils.NewUnexpectedInputError(data[endIdx])
					}
					endFileID = endIdx / 2
					endRemaining = int(data[endIdx] - '0')
				}
				totalValue += endFileID * positionidx
				positionidx++
				endRemaining--
			}
		}
	}

	return totalValue, nil
}
