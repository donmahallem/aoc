package day02

import (
	"github.com/donmahallem/aoc/go/aoc_utils"
)

type TOKEN int

const (
	TOKEN_STONE    = 1
	TOKEN_PAPER    = 2
	TOKEN_SCISSORS = 3
)

func mapToken(b byte) (int, error) {
	switch b {
	case 'A', 'X':
		return TOKEN_STONE, nil
	case 'B', 'Y':
		return TOKEN_PAPER, nil
	case 'C', 'Z':
		return TOKEN_SCISSORS, nil
	default:
		return 0, aoc_utils.NewUnexpectedInputError(b)
	}
}
