package day07

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func parseInput(in io.Reader) ([]int, error) {
	raw, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	counts := make([]int, 0, 1024)
	currentNumber := 0
	numberRead := false

	for _, b := range raw {
		switch b {
		case '\n', '\r':
			if numberRead {
				counts = append(counts, currentNumber)
				currentNumber = 0
				numberRead = false
			}
		case ',':
			if numberRead {
				counts = append(counts, currentNumber)
				currentNumber = 0
				numberRead = false
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			currentNumber = currentNumber*10 + int(b-'0')
			numberRead = true
		default:
			return nil, aoc_utils.NewUnexpectedInputError(b)
		}
	}

	if numberRead {
		counts = append(counts, currentNumber)
	}

	return counts, nil
}
