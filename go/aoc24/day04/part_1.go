package day04

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Dir aoc_utils.Point[int16]

var SearchTerm = []byte{'X', 'M', 'A', 'S'}
var SearchTermDirections = []Dir{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func CheckBlock(block [][]byte) int {
	height := int16(len(block))
	width := int16(len(block[0]))
	count := 0
	found := false
	for x := range width {
		for y := range height {
			for _, dir := range SearchTermDirections {
				endX := x + dir.X*3
				endY := y + dir.Y*3
				if endX < 0 || endY < 0 || endX >= width || endY >= height {
					continue
				}
				found = true
				for i := range int16(4) {
					if block[y+i*dir.Y][x+i*dir.X] != SearchTerm[i] {
						found = false
						break
					}
				}
				if found {
					count++
				}
			}
		}
	}
	return count
}
func Part1(in io.Reader) int {
	s := bufio.NewScanner(in)
	data := [][]byte{}
	for s.Scan() {
		lineData := s.Bytes()
		data = append(data, make([]byte, len(lineData)))
		copy(data[len(data)-1], lineData)
	}
	return CheckBlock(data)
}
