package day04

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type coord struct {
	row uint8
	col uint8
}
type bingoMap map[uint8]coord

func parseInput(in io.Reader) ([]uint8, []bingoMap, error) {
	numbers := make([]uint8, 0, 128)
	boards := make([]bingoMap, 0, 32)
	// expected comma seperated list of numbers
	s := bufio.NewScanner(in)
	if !s.Scan() {
		return nil, nil, s.Err()
	}
	line := s.Bytes()
	currentNumber := uint8(0)
	for i := 0; i < len(line); i++ {
		if line[i] == ',' {
			numbers = append(numbers, currentNumber)
			currentNumber = 0
		} else if line[i] >= '0' && line[i] <= '9' {
			currentNumber = currentNumber*10 + uint8(line[i]-'0')
		} else if line[i] == ' ' {
			// skip spaces
		} else {
			return nil, nil, aoc_utils.NewParseErrorf("Unexpected character in number line %c", line[i])
		}
	}
	// append the last number
	numbers = append(numbers, currentNumber)
	// parse the bingo boards
	currentLine := 0

	var currentBingoNumber uint8 = 0

	currentBingoMap := make(bingoMap)

	// always expects max 2 digit numbers
	// in column format "xx xx xx xx xx"
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			if currentLine%5 != 0 {
				return nil, nil, aoc_utils.NewParseErrorf("Incomplete bingo board")
			}
			continue
		}
		lineLen := len(line)
		if lineLen < 14 {
			return nil, nil, aoc_utils.NewParseErrorf("Unexpected line length in bingo board")
		}
		for i := 0; i < 14; {
			colIdx := i / 3
			if line[i] >= '0' && line[i] <= '9' {
				currentBingoNumber = currentBingoNumber*10 + uint8(line[i]-'0')
			}
			if i%3 == 2 {
				currentBingoMap[currentBingoNumber] = coord{uint8(currentLine), uint8(colIdx)}
				currentBingoNumber = 0
				i += 1
				continue
			}
			i++
		}
		// add line end number
		currentBingoMap[currentBingoNumber] = coord{uint8(currentLine), 4}
		currentLine++
		if currentLine < 5 {
			currentBingoNumber = 0
		} else if currentLine == 5 {
			boards = append(boards, currentBingoMap)
			currentBingoMap = make(bingoMap)
			currentLine = 0
			currentBingoNumber = 0
		}
	}

	return numbers, boards, nil
}

// represents hits in a 5x5 bingo board but counts only columns and rows [0-4 for rows, 5-9 for columns]
type luckyDraws [10]uint8
