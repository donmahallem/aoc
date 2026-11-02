package day04

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

// returns draw number to first bingo win
func playPart1(draws []uint8, boards []bingoMap) (int, error) {
	bingoPlays := make([]luckyDraws, len(boards))
	for _, draw := range draws {
		for bIdx, board := range boards {
			if coord, ok := board[draw]; ok {
				r, c := coord.row, coord.col
				bingoPlays[bIdx][r]++
				bingoPlays[bIdx][5+c]++
				delete(board, draw)
				if bingoPlays[bIdx][r] == 5 || bingoPlays[bIdx][5+c] == 5 {
					sumValues := 0
					for k := range board {
						sumValues += int(k)
					}
					return sumValues * int(draw), nil
				}
			}
		}
	}
	return -1, aoc_utils.NewSolverError("No bingo reached", nil)
}

func Part1(in io.Reader) (int, error) {
	draws, boards, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	// Implement the logic for Part1 using draws and boards
	return playPart1(draws, boards)
}
