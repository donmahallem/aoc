package day01

import (
	"bufio"
	"io"
)

type inputData struct {
	left     bool
	distance int
}

func parseInput(in io.Reader) []inputData {
	s := bufio.NewScanner(in)
	inputDataList := make([]inputData, 0, 300)
	for s.Scan() {
		inp := inputData{}
		line := s.Bytes()
		inp.left = line[0] == 'L'
		for i := 1; i < len(line); i++ {
			inp.distance = inp.distance*10 + int(line[i]-'0')
		}
		inputDataList = append(inputDataList, inp)
	}
	return inputDataList
}

func Part1(in io.Reader) int {
	data := parseInput(in)
	currentPosition := 50
	zeros := 0
	for _, d := range data {
		if d.left {
			currentPosition = (currentPosition - d.distance) % 100
		} else {
			currentPosition = (currentPosition + d.distance) % 100
		}
		if currentPosition == 0 {
			zeros++
		}
	}
	return zeros
}
