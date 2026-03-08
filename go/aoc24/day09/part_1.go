package day09

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

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
