package day07

import (
	"bufio"
	"io"
)

type splitterPositions map[int]struct{}

func parseInput(in io.Reader) (splitterPositions, int, int, int, int) {
	s := bufio.NewScanner(in)
	splitterMap := make(splitterPositions, 256)
	var startX, startY, height int
	width := -1
	for ; s.Scan(); height++ {
		line := s.Bytes()
		if width < 0 {
			width = len(line) // set width before using it
		}
		for x := range line {
			switch line[x] {
			case 'S':
				startX = x
				startY = height
			case '^':
				splitterMap[height*width+x] = struct{}{}
			}
		}
	}

	return splitterMap, startX, startY, width, height
}
