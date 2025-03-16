package day04

import (
	"bufio"
	"fmt"
	"os"
)

var SearchTermDirectionsMas = [][]int{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func CheckMasBlock(block [][]byte) int {
	height := len(block)
	width := len(block[0])
	count := 0
	totalCounter := 0
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if block[y][x] != 'A' {
				continue
			}
			count = 0
			for _, dir := range SearchTermDirectionsMas {
				xStart := x + dir[0]
				xEnd := x - dir[0]
				yStart := y + dir[1]
				yEnd := y - dir[1]
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
func Part2() {
	s := bufio.NewScanner(os.Stdin)
	data := [][]byte{}
	for s.Scan() {
		lineData := s.Bytes()
		data = append(data, make([]byte, len(lineData)))
		copy(data[len(data)-1], lineData)
	}
	result := CheckMasBlock(data)
	fmt.Printf("%d\n", result)
}
