package day02

import "github.com/donmahallem/aoc/go/aoc_utils/bytes"

type block struct {
	Red, Green, Blue int
}

func parseLine(line []byte) (int, []block) {
	game := 0
	curBlock := block{}
	blocks := make([]block, 0)
	curVal := 0
	for idx := 0; idx < len(line); idx++ {
		chr := line[idx]
		if val, ok := bytes.ParseIntFromByte[int](chr); ok {
			curVal = (curVal * 10) + val
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
			curBlock = block{}
		}
	}
	if curBlock.Blue+curBlock.Green+curBlock.Red > 0 {
		blocks = append(blocks, curBlock)
	}
	return game, blocks
}
