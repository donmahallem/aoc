package day04

import (
	"bufio"
	"io"
)

/*
checkMasBlock counts MAS crosses inside the three-row window
and the index of the row that represents the current top line.
*/
func checkMasBlock(block []byte, width, topIdx int) int {
	const windowHeight = 3
	if width < 3 || len(block) != windowHeight*width {
		return 0
	}
	topIdx = ((topIdx % windowHeight) + windowHeight) % windowHeight
	midIdx := (topIdx + 1) % windowHeight
	botIdx := (topIdx + 2) % windowHeight

	topStart := topIdx * width
	midStart := midIdx * width
	botStart := botIdx * width

	count := 0
	for x := 1; x < width-1; x++ {
		if block[midStart+x] != 'A' {
			continue
		}
		tl := block[topStart+x-1]
		tr := block[topStart+x+1]
		bl := block[botStart+x-1]
		br := block[botStart+x+1]

		if ((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) &&
			((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
			count++
		}
	}
	return count
}

func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	const windowHeight = 3
	var data []byte
	width := -1
	total := 0

	for rowsSeen := 0; s.Scan(); rowsSeen++ {
		line := s.Bytes()
		if width < 0 {
			width = len(line)
			if width < 3 {
				return 0
			}
			data = make([]byte, windowHeight*width)
		}
		if len(line) != width {
			panic("Inconsistent line widths")
		}

		idx := rowsSeen % windowHeight
		copy(data[idx*width:(idx+1)*width], line)

		if rowsSeen >= windowHeight-1 {
			topIdx := (idx + 1) % windowHeight
			total += checkMasBlock(data, width, topIdx)
		}

	}

	return total
}
