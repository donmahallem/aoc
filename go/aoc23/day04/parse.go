package day04

import (
	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

func parseLine(data []byte) (uint8, []uint8, []uint8) {
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
		if parsedInt, ok := bytes.ParseIntFromByte[uint8](data[idx]); ok {
			if !isInNumber {
				currentNumber = parsedInt
				isInNumber = true
			} else {
				currentNumber = (currentNumber * 10) + parsedInt
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
