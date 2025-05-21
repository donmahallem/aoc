package day02

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

func CalculateMinBlock(blocks *[]Block) int {
	block := Block{}
	for _, curBlock := range *blocks {
		block.Red = aoc_utils.Max(block.Red, curBlock.Red)
		block.Green = aoc_utils.Max(block.Green, curBlock.Green)
		block.Blue = aoc_utils.Max(block.Blue, curBlock.Blue)
	}
	return block.Red * block.Green * block.Blue
}

func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		d := s.Bytes()
		_, blocks := ParseLine(d)
		summe += CalculateMinBlock(&blocks)
	}
	return summe
}
