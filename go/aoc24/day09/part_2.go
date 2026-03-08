package day09

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part2(in io.Reader) (int, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return 0, err
	}
	for len(data) > 0 && (data[len(data)-1] == '\n' || data[len(data)-1] == '\r') {
		data = data[:len(data)-1]
	}
	if len(data) == 0 {
		return 0, nil
	}

	numFiles := (len(data) + 1) / 2
	// tracks block start positions
	blockPos := make([]int, numFiles)
	// tracks gap start positions
	gapStart := make([]int, numFiles)
	// tracks the space available for each gap
	gapAvail := make([]int, numFiles)

	pos := 0
	for i := range numFiles {
		idx := i * 2
		if data[idx] < '0' || data[idx] > '9' {
			return 0, aoc_utils.NewUnexpectedInputError(data[idx])
		}
		blockPos[i] = pos
		pos += int(data[idx] - '0')

		if gapIdx := idx + 1; gapIdx < len(data) {
			if data[gapIdx] < '0' || data[gapIdx] > '9' {
				return 0, aoc_utils.NewUnexpectedInputError(data[gapIdx])
			}
			sz := int(data[gapIdx] - '0')
			gapStart[i] = pos
			gapAvail[i] = sz
			pos += sz
		}
	}

	total := 0
	firstGap := [10]int{}

	for id := numFiles - 1; id >= 0; id-- {
		size := int(data[id*2] - '0')
		if size == 0 {
			continue
		}

		placed := false
		j := firstGap[size]
		for ; j < id; j++ {
			if gapAvail[j] >= size {
				startPos := gapStart[j]
				total += id * (size*startPos + size*(size-1)/2)
				gapStart[j] += size
				gapAvail[j] -= size
				placed = true
				break
			}
		}

		// All indices < j definitively have avail < size,
		// which means they also have avail < s for any s >= size.
		for s := size; s <= 9; s++ {
			if firstGap[s] < j {
				firstGap[s] = j
			}
		}

		if !placed {
			startPos := blockPos[id]
			total += id * (size*startPos + size*(size-1)/2)
		}
	}
	return total, nil
}
