package day18

import (
	"bufio"
	_ "embed"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

//go:embed sample.txt
var sample string

const (
	dirUp    = 'U'
	dirRight = 'R'
	dirDown  = 'D'
	dirLeft  = 'L'
)

type instruction struct {
	dir   uint8
	steps int64
}

type instructions = []instruction

func parseInput(r io.Reader, part1 bool) (*instructions, error) {
	scanner := bufio.NewScanner(r)
	instructionList := make(instructions, 0, 64)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 3 {
			return nil, aoc_utils.NewParseError("malformed instruction line", nil)
		}

		ins := instruction{dir: fields[0][0]}

		if part1 {
			steps, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				continue
			}
			ins.steps = steps
		} else {
			hexToken := fields[2]
			if len(hexToken) != 9 {
				return nil, aoc_utils.NewParseError("malformed hex token", nil)
			}
			steps, err := strconv.ParseInt(hexToken[2:7], 16, 64)
			if err != nil {
				return nil, aoc_utils.NewParseError("malformed hex token steps", err)
			}
			ins.steps = steps

			switch hexToken[7] {
			case '0':
				ins.dir = dirRight
			case '1':
				ins.dir = dirDown
			case '2':
				ins.dir = dirLeft
			case '3':
				ins.dir = dirUp
			default:
				continue
			}
		}

		instructionList = append(instructionList, ins)
	}
	return &instructionList, nil
}

/*
Using https://en.wikipedia.org/wiki/Shoelace_formula

Pick's theorem should work but due to the way the instructions are given
it is easier to use the shoelace formula.
*/
func calculateArea(ins *instructions) int64 {
	var totalArea int64
	var perimeter int64
	var currentX, currentY int64

	for _, inst := range *ins {
		nextX, nextY := currentX, currentY
		switch inst.dir {
		case dirUp:
			nextY += inst.steps
		case dirDown:
			nextY -= inst.steps
		case dirLeft:
			nextX -= inst.steps
		case dirRight:
			nextX += inst.steps
		}

		totalArea += currentX*nextY - nextX*currentY
		perimeter += inst.steps
		currentX, currentY = nextX, nextY
	}

	if totalArea < 0 {
		totalArea = -totalArea
	}

	return (totalArea / 2) + (perimeter / 2) + 1
}
