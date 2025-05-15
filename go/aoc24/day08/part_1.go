package day08

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Point aoc_utils.Point[int16]

func readSource(reader io.Reader) (map[byte][]Point, int16, int16) {
	data := make(map[byte][]Point, 0)
	s := bufio.NewScanner(reader)
	y := int16(0)
	width := int16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = int16(len(line))
		} else if width != int16(len(line)) {
			panic("Line length is uneven")
		}
		for idx, character := range line {
			if character == '.' {
				continue
			}
			if charData, ok := data[character]; ok {
				data[character] = append(charData, Point{Y: y, X: int16(idx)})
			} else {
				data[character] = []Point{{Y: y, X: int16(idx)}}
			}
		}
		y++
	}
	return data, width, y
}

func OutOfBounds(x int16, y int16, width int16, height int16) bool {
	return x < 0 || y < 0 || x >= width || y >= height
}

func Part1(in io.Reader) int {
	antennas, width, height := readSource(in)
	var antennaListLen int
	echos := make(map[[2]int16]bool, 0)
	for antenna := range antennas {
		antennaList := antennas[antenna]
		antennaListLen = len(antennaList)
		for i := range antennaListLen {
			for j := range antennaListLen {
				if i == j {
					continue
				}
				newX := 2*antennaList[i].X - antennaList[j].X
				newY := 2*antennaList[i].Y - antennaList[j].Y
				if !OutOfBounds(newX, newY, width, height) {
					echos[[2]int16{newY, newX}] = true
				}
			}
		}
	}
	return len(echos)
}
