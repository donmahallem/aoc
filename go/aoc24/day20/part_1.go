package day20

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Point = aoc_utils.Point[int]
type PointList = []Point

var DIR_UP Point = Point{X: 0, Y: -1}
var DIR_LEFT Point = Point{X: -1, Y: 0}
var DIR_RIGHT Point = Point{X: 1, Y: 0}
var DIR_DOWN Point = Point{X: 0, Y: 1}
var DIRS_ALL []*Point = []*Point{&DIR_DOWN, &DIR_UP, &DIR_LEFT, &DIR_RIGHT}

func ParseInput(in io.Reader) *PointList {
	var startPoint, endPoint Point
	pList := make([]Point, 0)
	s := bufio.NewScanner(in)
	y := 0
	for s.Scan() {
		line := s.Bytes()
		for x := range len(line) {
			switch line[x] {
			case '.':
				pList = append(pList, Point{X: x, Y: y})
			case 'S':
				startPoint = Point{X: x, Y: y}
			case 'E':
				endPoint = Point{X: x, Y: y}
			}
		}
		y++
	}
	// Append endPoint twice
	// Prepend start to list without array duplications
	pList = append(pList, endPoint, endPoint)
	copy(pList[1:], pList)
	pList[0] = startPoint
	// Iterate and find non visited neighbours
	for i := 1; i < len(pList)-2; i++ {
		for j := i + 1; j < len(pList)-1; j++ {
			if pList[i-1].DistanceManhatten(pList[j]) == 1 {
				pList[i], pList[j] = pList[j], pList[i]
				break
			}
		}
	}
	return &pList
}

func CountCheats(racewayPoints *[]Point, minSavings int) int {
	cheatCount := 0
	for leftIdx := range len(*racewayPoints) - minSavings {
		for rightIdx := len(*racewayPoints) - 1; leftIdx+minSavings < rightIdx; rightIdx-- {
			if (*racewayPoints)[leftIdx].DistanceManhatten((*racewayPoints)[rightIdx]) == 2 {
				cheatCount++
			}
		}
	}
	return cheatCount
}

func Part1(in io.Reader) int {
	patterns := ParseInput(in)
	return CountCheats(patterns, 100)
}
