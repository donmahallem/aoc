package day02

import (
	"bufio"
	"io"
)

type Block struct {
	Red, Green, Blue int
}

func ParseLine(line []byte) (int, []Block) {
	game := 0
	curBlock := Block{}
	blocks := make([]Block, 0)
	curVal := 0
	for idx := 0; idx < len(line); idx++ {
		chr := line[idx]
		if chr >= '0' && chr <= '9' {
			curVal = (curVal * 10) + int(chr-'0')
		} else if chr == ':' {
			game = curVal
			curVal = 0
		} else if chr == 'g' {
			curBlock.Green += curVal
			idx += 4
			curVal = 0
		} else if chr == 'b' {
			curBlock.Blue += curVal
			idx += 3
			curVal = 0
		} else if chr == 'r' {
			curBlock.Red += curVal
			idx += 2
			curVal = 0
		} else if chr == ';' {
			blocks = append(blocks, curBlock)
			curBlock = Block{}
		}
	}
	if curBlock.Blue+curBlock.Green+curBlock.Red > 0 {
		blocks = append(blocks, curBlock)
	}
	return game, blocks
}

func ValidateBlocks(blocks []Block) bool {
	for _, block := range blocks {
		if block.Red > 12 || block.Green > 13 || block.Blue > 14 {
			return false
		}
	}
	return true
}

func Part1(in io.Reader) int {
	s := bufio.NewScanner(in)
	summe := 0
	for s.Scan() {
		d := s.Bytes()
		gameId, blocks := ParseLine(d)
		if ValidateBlocks(blocks) {
			summe += gameId
		}
	}
	return summe
}
