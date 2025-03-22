package day09

import (
	"fmt"
	"io"
	"os"
)

func FindEmptySpace(inp *[]int16, size *int, rangeEnd *int) (int, int, bool) {
	var startIdx, endIdx int
	inside := false
	for i := range *rangeEnd {
		if !inside && (*inp)[i] == -1 {
			startIdx = i
			endIdx = startIdx + 1
			inside = true
		} else if inside {
			if (*inp)[i] == -1 {
				endIdx = i + 1
			} else {
				inside = false
			}
		}
		if inside && endIdx-startIdx == *size {
			return startIdx, endIdx, true
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
	blockId := (*inp)[len((*inp))-1]
	var lastBlockStart int = -1
	for ; blockId >= 0; blockId-- {
		if blockStart, blockEnd, blockOk := FindBlock(inp, &blockId, &lastBlockStart); blockOk {
			blockSize := blockEnd - blockStart
			if spaceStart, _, spaceOk := FindEmptySpace(inp, &blockSize, &blockStart); spaceOk {
				for i := 0; i < blockSize; i++ {
					(*inp)[spaceStart+i] = (*inp)[blockStart+i]
					(*inp)[blockStart+i] = -1
				}
			}
			lastBlockStart = blockStart
		}
	}
}

func Part2(in *os.File) {
	data, _ := io.ReadAll(in)
	expandedData := ConvertInput(&data)
	CompactLess(&expandedData)
	fmt.Printf("%d\n", CheckSum(&expandedData))
}
