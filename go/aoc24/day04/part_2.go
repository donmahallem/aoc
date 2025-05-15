package day04

import (
	"bufio"
	"io"
)

var SearchTermDirectionsMas = []Dir{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func CheckMasBlock(block [][]byte) int {
	height := int16(len(block))
	width := int16(len(block[0]))
	count := 0
	totalCounter := 0
	for x := int16(1); x < width-1; x++ {
		for y := int16(1); y < height-1; y++ {
			if block[y][x] != 'A' {
				continue
			}
			count = 0
			for _, dir := range SearchTermDirectionsMas {
				xStart := x + dir.X
				xEnd := x - dir.X
				yStart := y + dir.Y
				yEnd := y - dir.Y
				if block[yStart][xStart] == 'M' && block[yEnd][xEnd] == 'S' {
					count++
					if count >= 2 {
						break
					}
				}
			}
			if count >= 2 {
				totalCounter++
			}
		}
	}
	return totalCounter
}
func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	data := [][]byte{}
	for s.Scan() {
		lineData := s.Bytes()
		data = append(data, make([]byte, len(lineData)))
		copy(data[len(data)-1], lineData)
	}
	return CheckMasBlock(data)
}
