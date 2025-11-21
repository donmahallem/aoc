package day24

import (
	"bufio"
	"io"
)

func parseInput[A int64 | float64](r io.Reader) []hail[A] {
	hailData := make([]hail[A], 0, 300)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var h hail[A]
		line := scanner.Bytes()
		currentNum := &h.Px
		numNegative := false
		numStarted := false
		// 20, 19, 15 @  1, -5, -3
		for _, b := range line {
			switch b {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				numStarted = true
				*currentNum = *currentNum*10 + A(b-'0')
			case '-':
				numNegative = true
			case ',', ' ', '@':
				if numStarted {
					if numNegative {
						*currentNum = -*currentNum
					}
					// move to next number
					switch currentNum {
					case &h.Px:
						currentNum = &h.Py
					case &h.Py:
						currentNum = &h.Pz
					case &h.Pz:
						currentNum = &h.Vx
					case &h.Vx:
						currentNum = &h.Vy
					case &h.Vy:
						currentNum = &h.Vz
					}
					numNegative = false
					numStarted = false
				}
			}
		}
		if numNegative {
			*currentNum = -*currentNum
		}
		hailData = append(hailData, h)
	}
	return hailData
}
