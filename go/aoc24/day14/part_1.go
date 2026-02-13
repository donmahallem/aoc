package day14

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Robot struct {
	pos, vec aoc_utils.Point[int]
}

func NewRobot(posX, posY, vecX, vecY int) *Robot {
	return &Robot{*aoc_utils.NewPoint(posX, posY), *aoc_utils.NewPoint(vecX, vecY)}
}

type parsePosition uint8

const (
	parsePositionInitial parsePosition = iota
	parsePositionP1
	parsePositionP2
	parsePositionV1
	parsePositionV2
	parsePositionEq1
	parsePositionEq2
	parsePositionV
)

func ParseLine(line []byte) (*Robot, error) {
	// expected format: p=9,3 v=2,3
	robot := Robot{}

	readInt := func(startidx int) (val int, nextIdx int) {
		isNegative := false
		val = 0
		for i := startidx; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				val = (val * 10) + int(c-'0')
			} else if c == '-' {
				isNegative = true
			} else {
				if isNegative {
					val *= -1
				}
				return val, i
			}
		}
		if isNegative {
			val *= -1
		}
		return val, len(line)
	}
	if len(line) < 7 {
		return nil, aoc_utils.NewParseError("The input must be atleast 7 characters long", nil)
	}

	if line[0] != 'p' {
		return nil, aoc_utils.NewUnexpectedInputError(line[0])
	} else if line[1] != '=' {
		return nil, aoc_utils.NewUnexpectedInputError(line[1])
	}

	startIdx := 2
	robot.pos.X, startIdx = readInt(startIdx)
	if startIdx >= len(line) || line[startIdx] != ',' {
		return nil, aoc_utils.NewParseError("expected ',' after pos.X", nil)
	}
	startIdx++
	robot.pos.Y, startIdx = readInt(startIdx)
	if startIdx >= len(line) || line[startIdx] != ' ' {
		return nil, aoc_utils.NewParseError("expected ' ' after pos.Y", nil)
	}
	startIdx++
	if startIdx >= len(line) || line[startIdx] != 'v' {
		return nil, aoc_utils.NewParseError("expected 'v' after position", nil)
	}
	startIdx++
	if startIdx >= len(line) || line[startIdx] != '=' {
		return nil, aoc_utils.NewParseError("expected '=' after 'v'", nil)
	}
	startIdx++
	robot.vec.X, startIdx = readInt(startIdx)
	if startIdx >= len(line) || line[startIdx] != ',' {
		return nil, aoc_utils.NewParseError("expected ',' after vec.X", nil)
	}
	startIdx++
	robot.vec.Y, startIdx = readInt(startIdx)

	return &robot, nil
}

func LoadFile(reader io.Reader) ([]Robot, error) {
	obstacles := make([]Robot, 0, 100)
	s := bufio.NewScanner(reader)
	for s.Scan() {
		line := s.Bytes()
		robot, err := ParseLine(line)
		if err != nil {
			return nil, err
		}
		obstacles = append(obstacles, *robot)
	}
	return obstacles, nil
}

func CalculateQuadrant(robot *Robot, steps int, width int, height int) int8 {
	finalX := (robot.pos.X + (robot.vec.X * steps)) % width
	finalY := (robot.pos.Y + (robot.vec.Y * steps)) % height
	if finalX < 0 {
		finalX = width + finalX
	}
	if finalY < 0 {
		finalY = height + finalY
	}
	dividerX := width / 2
	dividerY := height / 2
	if finalX == dividerX || finalY == dividerY {
		return -1
	}
	var val int8 = 0
	if finalX > dividerX {
		val = 1
	}
	if finalY > dividerY {
		val += 2
	}
	return val
}

func CountQuadrant(robots []Robot, steps int, width int, height int) int {
	quadrantMap := [4]int{}
	for _, robot := range robots {
		quadrant := CalculateQuadrant(&robot, steps, width, height)
		if quadrant >= 0 {
			quadrantMap[quadrant]++
		}
	}
	countSum := 1
	for key := range 4 {
		if quadrantMap[key] == 0 {
			continue
		}
		countSum *= quadrantMap[key]
	}
	return countSum
}

func Part1(in io.Reader) (int, error) {
	robots, err := LoadFile(in)
	if err != nil {
		return 0, err
	}
	maxX, maxY := 0, 0
	for _, r := range robots {
		if r.pos.X > maxX {
			maxX = r.pos.X
		}
		if r.pos.Y > maxY {
			maxY = r.pos.Y
		}
	}
	width := maxX + 1
	height := maxY + 1
	if width <= 0 || height <= 0 {
		return 0, aoc_utils.NewParseError("invalid grid size", nil)
	}
	totalSum := CountQuadrant(robots, 100, width, height)
	return totalSum, nil
}
