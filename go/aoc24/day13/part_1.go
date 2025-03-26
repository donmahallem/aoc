package day13

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Input struct {
	vec1, vec2 [2]float64
	target     [2]float64
}

func ParseButton(data []byte) [2]float64 {
	var ret [2]float64
	first := true
	var val uint8
	for i := range len(data) {
		val = data[i] - '0'
		if val >= 0 && val <= 9 {
			if first {
				ret[0] = (ret[0] * 10) + float64(val)
			} else {
				ret[1] = (ret[1] * 10) + float64(val)
			}
		} else if data[i] == 'Y' {
			first = false
		}
	}
	return ret
}
func ParseResult(data []byte) [2]float64 {
	var ret [2]float64
	first := true
	var val uint8
	for i := range len(data) {
		val = data[i] - '0'
		if val >= 0 && val <= 9 {
			if first {
				ret[0] = (ret[0] * 10) + float64(val)
			} else {
				ret[1] = (ret[1] * 10) + float64(val)
			}
		} else if data[i] == 'Y' {
			first = false
		}
	}
	return ret
}
func LoadFile(reader io.Reader) []Input {
	obstacles := make([]Input, 0, 100)
	s := bufio.NewScanner(reader)
	var currentInput Input = Input{}
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			obstacles = append(obstacles, currentInput)
			currentInput = Input{}
			continue
		}
		switch line[7] {
		case 'A':
			currentInput.vec1 = ParseButton(line[7:])
			break
		case 'B':
			currentInput.vec2 = ParseButton(line[7:])
			break
		case 'X':
			currentInput.target = ParseButton(line[6:])
			break
		}
	}
	obstacles = append(obstacles, currentInput)
	return obstacles
}

func Calculate(inp *Input) (int, int, bool) {
	vec2_factor := (((*inp).target[1] * (*inp).vec1[0]) - ((*inp).vec1[1] * (*inp).target[0])) / (((*inp).vec2[1] * (*inp).vec1[0]) - ((*inp).vec1[1] * (*inp).vec2[0]))
	vec1_factor := ((*inp).target[1] - (vec2_factor * (*inp).vec2[1])) / (*inp).vec1[1]
	if math.Trunc(vec1_factor) == vec1_factor && math.Trunc(vec2_factor) == vec2_factor {
		return int(vec1_factor), int(vec2_factor), true
	}
	return 0, 0, false
}

func Part1(in *os.File) {
	data := LoadFile(in)
	totalSum := 0
	for _, inp := range data {
		a, b, ok := Calculate(&inp)
		if !ok {
			continue
		}
		totalSum += a*3 + b
	}
	fmt.Printf("%d\n", totalSum)
}
