package day02

import (
	"bufio"
	"fmt"
	"io"
	"os"

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

func ParseFilePart2(reader io.Reader) int {
	s := bufio.NewScanner(reader)
	summe := 0
	for s.Scan() {
		d := s.Bytes()
		_, blocks := ParseLine(&d)
		summe += CalculateMinBlock(&blocks)
	}
	return summe
}
func Part2(in *os.File) {
	fmt.Printf("Result: %d\n", ParseFilePart2(in))
}
