package day13

import (
	"bufio"
	"io"
	"math"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type VecFloat64 = aoc_utils.Point[float64]
type Input struct {
	vec1, vec2 VecFloat64
	target     VecFloat64
}

func ParseButton(data []byte) VecFloat64 {
	var ret VecFloat64
	first := true
	for i := 0; i < len(data); i++ {
		ch := data[i]
		if ch >= '0' && ch <= '9' {
			val := float64(ch - '0')
			if first {
				ret.Y = (ret.Y * 10) + val
			} else {
				ret.X = (ret.X * 10) + val
			}
		} else if ch == 'Y' {
			first = false
		}
	}
	return ret
}

func LoadFile(reader io.Reader) ([]Input, error) {
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
		if len(line) <= 7 {
			return nil, aoc_utils.NewParseError("malformed line", nil)
		}
		switch line[7] {
		case 'A':
			currentInput.vec1 = ParseButton(line[7:])
		case 'B':
			currentInput.vec2 = ParseButton(line[7:])
		case 'X':
			currentInput.target = ParseButton(line[7:])
		}
	}
	obstacles = append(obstacles, currentInput)
	return obstacles, nil
}

func Calculate(inp *Input) (int, int, bool) {
	vec2_factor := (((*inp).target.X * (*inp).vec1.Y) - ((*inp).vec1.X * (*inp).target.Y)) / (((*inp).vec2.X * (*inp).vec1.Y) - ((*inp).vec1.X * (*inp).vec2.Y))
	vec1_factor := ((*inp).target.X - (vec2_factor * (*inp).vec2.X)) / (*inp).vec1.X
	if math.Trunc(vec1_factor) == vec1_factor && math.Trunc(vec2_factor) == vec2_factor {
		return int(vec1_factor), int(vec2_factor), true
	}
	return 0, 0, false
}

func Part1(in io.Reader) (int, error) {
	data, err := LoadFile(in)
	if err != nil {
		return 0, err
	}
	totalSum := 0
	for _, inp := range data {
		a, b, ok := Calculate(&inp)
		if !ok {
			continue
		}
		totalSum += a*3 + b
	}
	return totalSum, nil
}
