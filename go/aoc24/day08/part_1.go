package day08

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readSource(reader io.Reader) (map[byte][][2]int16, int16, int16) {
	data := make(map[byte][][2]int16, 0)
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
				data[character] = append(charData, [2]int16{y, int16(idx)})
			} else {
				data[character] = [][2]int16{{y, int16(idx)}}
			}
		}
		y++
	}
	return data, width, y
}

func OutOfBounds(x int16, y int16, width int16, height int16) bool {
	return x < 0 || y < 0 || x >= width || y >= height
}

func Part1() {
	antennas, width, height := readSource(os.Stdin)
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
				newX := 2*antennaList[i][1] - antennaList[j][1]
				newY := 2*antennaList[i][0] - antennaList[j][0]
				if !OutOfBounds(newX, newY, width, height) {
					echos[[2]int16{newY, newX}] = true
				}
			}
		}
	}
	fmt.Printf("%d\n", len(echos))
}
