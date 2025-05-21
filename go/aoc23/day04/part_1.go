package day04

import (
	"bufio"
	"io"
	"slices"
)

func IsNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func ParseLine(data []byte) (uint8, []uint8, []uint8) {
	ticketFound := false
	isInNumber := false
	var ticket uint8 = 0
	winningBlock := true
	winningNumbers := make([]uint8, 0, 10)
	pickedNumbers := make([]uint8, 0, 25)
	var currentNumber uint8 = 0
	finishNum := func() {
		if isInNumber {
			if !ticketFound {
				ticket = currentNumber
				ticketFound = true
			} else if winningBlock {
				winningNumbers = append(winningNumbers, currentNumber)
			} else {
				pickedNumbers = append(pickedNumbers, currentNumber)
			}
			isInNumber = false
		}
	}
	for idx := 4; idx < len(data); idx++ {
		if IsNumber(data[idx]) {
			if !isInNumber {
				currentNumber = data[idx] - '0'
				isInNumber = true
			} else {
				currentNumber = (currentNumber * 10) + data[idx] - '0'
			}
			continue
		}
		finishNum()
		if data[idx] == '|' {
			winningBlock = false
		}
	}
	finishNum()
	return ticket, winningNumbers, pickedNumbers
}

func CountWinnings(winners []uint8, picks []uint8) int {
	score := 0
	for _, winner := range winners {
		if slices.Contains(picks, winner) {
			score++
		}
	}
	return score
}

func GetScore(winners []uint8, picks []uint8) int {
	score := CountWinnings(winners, picks)
	if score > 0 {
		return 1 << (score - 1)
	}
	return 0
}

func Part1(in io.Reader) int {
	s := bufio.NewScanner(in)
	score := 0
	for s.Scan() {
		_, a, b := ParseLine(s.Bytes())
		score += GetScore(a, b)
	}
	return score
}
