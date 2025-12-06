package day04

import (
	"bufio"
	"io"
)

type field []byte

func parseInput(in io.Reader) (field, int, int) {
	s := bufio.NewScanner(in)
	var g field
	rows := 0
	width := 0

	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = len(line)
			g = make(field, 0, width*width)
		}
		for i := range len(line) {
			if line[i] == '@' {
				g = append(g, 1)
			} else {
				g = append(g, 0)
			}
		}
		rows++
	}
	return g, width, rows
}
