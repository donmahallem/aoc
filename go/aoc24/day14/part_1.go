package day14

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type PointType interface {
	int | int8 | float32 | float64
}
type Point[A PointType] struct {
	x, y A
}

func (a *Point[A]) diff(b Point[A]) *Point[A] {
	return NewPoint(b.x-a.x, b.y-a.y)
}

func NewPoint[A PointType](x, y A) *Point[A] {
	return &Point[A]{x, y}
}

type Robot struct {
	pos, vec Point[int]
}

func NewRobot(posX, posY, vecX, vecY int) *Robot {
	return &Robot{*NewPoint(posX, posY), *NewPoint(vecX, vecY)}
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
	currentPointer := &robot.pos.x
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
			case &robot.pos.x:
				currentPointer = &robot.pos.y
				break
			case &robot.pos.y:
				currentPointer = &robot.vec.x
				break
			case &robot.vec.x:
				currentPointer = &robot.vec.y
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
	finalX := (robot.pos.x + (robot.vec.x * steps)) % width
	finalY := (robot.pos.y + (robot.vec.y * steps)) % height
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

func Part1(in *os.File) {
	data := LoadFile(in)
	totalSum := CountQuadrant(&data, 100, 101, 103)
	fmt.Printf("%d\n", totalSum)
}
