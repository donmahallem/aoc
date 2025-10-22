package day13

import (
	"bufio"
	"io"
)

type Block struct {
	Rows []int
	Cols []int
}

func ParseInput(r io.Reader) []Block {
	scanner := bufio.NewScanner(r)

	blocks := make([]Block, 0, 4)
	block := Block{
		Rows: make([]int, 0, 16),
		Cols: nil,
	}

	currentHeight := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			currentHeight = 0
			if len(block.Rows) > 0 {
				blocks = append(blocks, block)
			}
			block = Block{
				Rows: make([]int, 0, 16),
				Cols: nil,
			}
			continue
		}
		if block.Cols == nil {
			block.Cols = make([]int, len(line))
		}
		currentLine := 0
		for idx, c := range line {

			if c == '#' {
				currentLine |= 1 << idx
				block.Cols[idx] |= 1 << currentHeight
			}
		}
		block.Rows = append(block.Rows, currentLine)
		currentHeight++
	}
	if currentHeight > 0 {
		blocks = append(blocks, block)
	}
	return blocks
}

func validateAxis(items []int, center int) bool {
	for offset := 1; center-offset-1 >= 0 && center+offset < len(items); offset++ {
		if items[center-offset-1] != items[center+offset] {
			return false
		}
	}
	return true
}

func findAxis(axisData []int) (int, bool) {
	for centerIdx := 1; centerIdx < len(axisData); centerIdx++ {
		if axisData[centerIdx-1] == axisData[centerIdx] {
			if validateAxis(axisData, centerIdx) {
				return centerIdx, true
			}
		}
	}
	return -1, false
}

func Part1(in io.Reader) int {
	start := ParseInput(in)
	accum := 0
	for _, block := range start {
		if rowAxis, ok := findAxis(block.Rows); ok {
			accum += rowAxis * 100
			continue
		}
		if colAxis, ok := findAxis(block.Cols); ok {
			accum += colAxis
			continue
		}
	}
	return accum
}
