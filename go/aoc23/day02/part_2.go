package day02

import (
	"bufio"
	"io"
)

func CalculateMinBlock(blocks *[]Block) int {
	block := Block{}
	for _, curBlock := range *blocks {
		block.Red = max(block.Red, curBlock.Red)
		block.Green = max(block.Green, curBlock.Green)
		block.Blue = max(block.Blue, curBlock.Blue)
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
