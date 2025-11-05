package day04

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Dir aoc_utils.Point[int16]

var searchTerm = []byte{'X', 'M', 'A', 'S'}

// Checks the first four rows of the block for occurrences of the search term in all 6 directions except horizontal.
func checkBlock(block []byte, width int) int {
	if width <= 0 || len(block)%width != 0 {
		return 0
	}
	searchLen := len(searchTerm)
	directions := [][2]int{{1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	count := 0
	for row := range 4 {
		for col := range width {
			if block[row*width+col] != searchTerm[0] {
				continue
			}
			for _, dir := range directions {
				endRow := row + dir[0]*(searchLen-1)
				endCol := col + dir[1]*(searchLen-1)
				if endRow < 0 || endRow >= 4 || endCol < 0 || endCol >= width {
					continue
				}
				match := true
				for i := 1; i < searchLen; i++ {
					r := row + dir[0]*i
					c := col + dir[1]*i
					if block[r*width+c] != searchTerm[i] {
						match = false
						break
					}
				}
				if match {
					count++
				}
			}
		}
	}
	return count
}

const searchTermLength = 4
const hashForward uint32 = 'X'<<24 | 'M'<<16 | 'A'<<8 | 'S'
const hashBackward uint32 = 'S'<<24 | 'A'<<16 | 'M'<<8 | 'X'

// CheckLine counts forward and backward horizontal occurrences of the search term.
func CheckLine(line []byte, width int) int {
	if width < searchTermLength {
		// nothing to check
		return 0
	}
	count := 0
	var runningHash uint32
	for i := range width {
		runningHash = (runningHash << 8) | uint32(line[i])
		if i+1 >= searchTermLength && (runningHash == hashForward || runningHash == hashBackward) {
			count++
		}
	}
	return count
}

func Part1(in io.Reader) int {
	s := bufio.NewScanner(in)
	var data []byte = nil
	width := -1
	count := 0
	for lines := 0; s.Scan(); lines++ {
		lineData := s.Bytes()
		if width < 0 {
			width = len(lineData)
			data = make([]byte, 4*width)
		}
		count += CheckLine(lineData, width)
		if lines < 4 {
			copy(data[lines*width:], lineData)
		} else {
			copy(data, data[width:])
			copy(data[3*width:], lineData)
		}
		if lines >= 3 {
			count += checkBlock(data, width)
		}
	}
	return count
}
