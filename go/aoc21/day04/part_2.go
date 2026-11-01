package day04

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func playPart2(draws []uint8, boards []bingoMap) (int, error) {
	bingoPlays := make([]luckyDraws, len(boards))
	boardsWon := make([]bool, len(boards))
	boardsWonCount := 0
	for _, draw := range draws {
		for bIdx, board := range boards {
			if boardsWon[bIdx] {
				continue
			}
			if coord, ok := board[draw]; ok {
				r, c := coord.row, coord.col
				bingoPlays[bIdx][r]++
				bingoPlays[bIdx][5+c]++
				delete(board, draw)
				if bingoPlays[bIdx][r] == 5 || bingoPlays[bIdx][5+c] == 5 {
					boardsWon[bIdx] = true
					boardsWonCount++
					if boardsWonCount == len(boards) {
						sumValues := 0
						for k := range board {
							sumValues += int(k)
						}
						return sumValues * int(draw), nil
					}
				}
			}
		}
	}
	return -1, aoc_utils.NewSolverError("No bingo reached", nil)
}
func Part2(in io.Reader) (int, error) {
	draws, boards, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	return playPart2(draws, boards)
}
