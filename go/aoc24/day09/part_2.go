package day09

import (
	"io"
)

func FindEmptySpace(inp *[]int16, size *int, rangeEnd *int) (int, int, bool) {
	n := len(*inp)
	limit := *rangeEnd
	if limit < 0 || limit > n {
		limit = n
	}
	startIdx := -1
	count := 0
	for i := 0; i < limit; i++ {
		if (*inp)[i] == -1 {
			if startIdx < 0 {
				startIdx = i
				count = 1
			} else {
				count++
			}
			if count == *size {
				return startIdx, i + 1, true
			}
		} else {
			startIdx = -1
			count = 0
		}
	}
	return -1, -1, false
}

func FindBlock(inp *[]int16, id *int16, maxIdx *int) (int, int, bool) {
	var startIdx, endIdx int
	inside := false
	searchStart := *maxIdx
	if searchStart < 0 {
		searchStart = len((*inp)) - 1
	}
	for i := searchStart; i >= 0; i-- {
		if !inside && (*inp)[i] == *id {
			startIdx = i
			endIdx = startIdx + 1
			inside = true
		} else if inside {
			if (*inp)[i] == *id {
				startIdx--
			} else {
				return startIdx, endIdx, true
			}
		}
	}
	return -1, -1, false
}

func CompactLess(inp *[]int16) {
	if len(*inp) == 0 {
		return
	}
	blockId := (*inp)[len((*inp))-1]
	var lastBlockStart int = -1
	for ; blockId >= 0; blockId-- {
		if blockStart, blockEnd, blockOk := FindBlock(inp, &blockId, &lastBlockStart); blockOk {
			blockSize := blockEnd - blockStart
			if blockSize <= 0 {
				lastBlockStart = blockStart
				continue
			}
			if spaceStart, _, spaceOk := FindEmptySpace(inp, &blockSize, &blockStart); spaceOk {
				// validate bounds
				n := len(*inp)
				if spaceStart < 0 || blockStart < 0 || spaceStart+blockSize > n || blockStart+blockSize > n {
					lastBlockStart = blockStart
					continue
				}
				for i := 0; i < blockSize; i++ {
					(*inp)[spaceStart+i] = (*inp)[blockStart+i]
					(*inp)[blockStart+i] = -1
				}
			}
			lastBlockStart = blockStart
		}
	}
}

func Part2(in io.Reader) (int, error) {
	data, _ := io.ReadAll(in)
	expandedData, err := convertInput(data)
	if err != nil {
		return 0, err
	}
	CompactLess(&expandedData)
	return checkSum(&expandedData), nil
}
