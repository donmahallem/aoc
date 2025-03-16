package day04

import (
	"bufio"
	"fmt"
	"os"
)

var SearchTerm = []byte{'X', 'M', 'A', 'S'}
var SearchTermDirections = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func CheckBlock(block [][]byte) int {
	height := len(block)
	width := len(block[0])
	count := 0
	found := false
	for x := range width {
		for y := range height {
			for _, dir := range SearchTermDirections {
				endX := x + dir[0]*3
				endY := y + dir[1]*3
				if endX < 0 || endY < 0 || endX >= width || endY >= height {
					continue
				}
				found = true
				for i := range 4 {
					if block[y+i*dir[1]][x+i*dir[0]] != SearchTerm[i] {
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
func Part1() {
	s := bufio.NewScanner(os.Stdin)
	data := [][]byte{}
	for s.Scan() {
		lineData := s.Bytes()
		data = append(data, make([]byte, len(lineData)))
		copy(data[len(data)-1], lineData)
	}
	result := CheckBlock(data)
	fmt.Printf("%d\n", result)
}
