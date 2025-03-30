package day04

import (
	"bufio"
	"io"
	"slices"
)

func IsNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func ParseLine(data []byte) (int, *[]int, *[]int) {
	ticket := -1
	winningBlock := true
	winningNumbers := make([]int, 0)
	pickedNumbers := make([]int, 0)
	currentNumber := -1
	finishNum := func() {
		if currentNumber >= 0 {
			if ticket < 0 {
				ticket = currentNumber
			} else if winningBlock {
				winningNumbers = append(winningNumbers, currentNumber)
			} else {
				pickedNumbers = append(pickedNumbers, currentNumber)
			}
			currentNumber = -1
		}
	}
	for idx := 4; idx < len(data); idx++ {
		if IsNumber(data[idx]) {
			if currentNumber < 0 {
				currentNumber = int(data[idx] - '0')
			} else {
				currentNumber = (currentNumber * 10) + int(data[idx]-'0')
			}
			continue
		}
		finishNum()
		if data[idx] == '|' {
			winningBlock = false
		}
	}
	finishNum()
	return ticket, &winningNumbers, &pickedNumbers
}

func CountWinnings(winners *[]int, picks *[]int) int {
	score := 0
	for _, winner := range *winners {
		if slices.Contains(*picks, winner) {
			score++
		}
	}
	return score
}

func GetScore(winners *[]int, picks *[]int) int {
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
