package day21

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Point = aoc_utils.Point[int8]

func ParseInput(in io.Reader) *[][]Point {
	data := make([][]Point, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Bytes()
		data = append(data, make([]Point, 0, len(line)))
		for x := range len(line) {
			data[len(data)-1] = append(data[len(data)-1], *translateChar(&line[x]))
		}
	}
	return &data
}
