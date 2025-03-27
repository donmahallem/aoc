package day14

import (
	"bufio"
	"fmt"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Robot struct {
	pos, vec aoc_utils.Point[int]
}

func NewRobot(posX, posY, vecX, vecY int) *Robot {
	return &Robot{*aoc_utils.NewPoint(posX, posY), *aoc_utils.NewPoint(vecX, vecY)}
}

type ParsePosition uint8

const (
	ParsePosition_X1 ParsePosition = iota
	ParsePosition_X2
	ParsePosition_Y1
	ParsePosition_Y2
)

func ParseLine(line *[]byte) Robot {
	robot := Robot{}
	var shiftVal int
	// Pointer to the current selected output field in robot
	currentPointer := &robot.pos.X
	isNegative := false
	for i := range len(*line) {
		if (*line)[i] >= '0' && (*line)[i] <= '9' {
			shiftVal = int((*line)[i] - '0')
			if *currentPointer == 0 {
				*currentPointer = shiftVal
			} else {
				*currentPointer = (*currentPointer * 10) + shiftVal
			}
		} else if (*line)[i] == '-' {
			isNegative = true
		} else if (*line)[i] == 'v' || (*line)[i] == ',' {
			if isNegative {
				*currentPointer *= -1
			}
			switch currentPointer {
			case &robot.pos.X:
				currentPointer = &robot.pos.Y
				break
			case &robot.pos.Y:
				currentPointer = &robot.vec.X
				break
			case &robot.vec.X:
				currentPointer = &robot.vec.Y
				break
			}
			isNegative = false
		}
	}
	if isNegative {
		*currentPointer *= -1
	}
	//p=9,3 v=2,3
	return robot
}

func LoadFile(reader io.Reader) []Robot {
	obstacles := make([]Robot, 0, 100)
	s := bufio.NewScanner(reader)
	for s.Scan() {
		line := s.Bytes()
		obstacles = append(obstacles, ParseLine(&line))
	}
	return obstacles
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

func CountQuadrant(robots *[]Robot, steps int, width int, height int) int {
	quadrantMap := [4]int{}
	for _, robot := range *robots {
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

func Part1(in io.Reader) {
	data := LoadFile(in)
	totalSum := CountQuadrant(&data, 100, 101, 103)
	fmt.Printf("%d\n", totalSum)
}
