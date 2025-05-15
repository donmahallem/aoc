package day18

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils"
)

const CELL_CORRUPTED int = -1

// As you always walk top-left to right-bottom primarly use those first
var DIRS_ALL [4]Point = [4]Point{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

type Field = [][]uint16
type Point = aoc_utils.Point[int16]

func ParseInput(in io.Reader) []Point {
	parsedPoints := make([]Point, 0, 64)
	s := bufio.NewScanner(in)
	i := uint16(1)
	for ; s.Scan(); i++ {
		line := s.Text()
		commaIndex := strings.Index(line, ",")
		x, _ := strconv.Atoi(line[0:commaIndex])
		y, _ := strconv.Atoi(line[commaIndex+1:])
		point := Point{X: int16(x), Y: int16(y)}
		parsedPoints = append(parsedPoints, point)
	}
	return parsedPoints
}

func PointsToField(points []Point, width uint8, height uint8) Field {

	parsedField := make(Field, height)
	for y := range height {
		parsedField[y] = make([]uint16, width)
	}
	var cellValue uint16
	for i, point := range points {
		cellValue = uint16(i) + 1
		if parsedField[point.Y][point.X] > 0 && parsedField[point.Y][point.X] < cellValue {
			continue
		}
		parsedField[point.Y][point.X] = cellValue
	}
	return parsedField
}

type PathNode struct {
	coord Point
	len   uint16
}

func FindShortestPath(field Field, stepsTaken uint16, fieldWidth int16, fieldHeight int16) uint16 {
	queue := make([]PathNode, 0, 64)
	queue = append(queue, PathNode{coord: Point{X: 0, Y: 0}, len: 0})
	visited := make(map[Point]bool, 128)
	var currentNode PathNode
	var shortestPath uint16 = math.MaxUint16
	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		if currentNode.len+1 >= shortestPath {
			continue
		}

		visited[currentNode.coord] = true
		for _, checkDir := range DIRS_ALL {
			nextPoint := Point{X: currentNode.coord.X + checkDir.X, Y: currentNode.coord.Y + checkDir.Y}

			if visited[currentNode.coord] {
				continue
			}
			nextNode := PathNode{coord: nextPoint}
			nextNode.len = currentNode.len + 1
			if nextNode.coord.X < 0 || nextNode.coord.Y < 0 || nextNode.coord.X >= fieldWidth || nextNode.coord.Y >= fieldHeight {
				// next coord outside the field dimensions
				continue
			}
			currentCellValue := field[nextNode.coord.Y][nextNode.coord.X]
			if currentCellValue > 0 && currentCellValue <= stepsTaken {
				// Cell Corrupted
				continue
			}
			if nextNode.coord.X == fieldWidth-1 && nextNode.coord.Y == fieldHeight-1 {
				shortestPath = nextNode.len
			} else {
				queue = append(queue, nextNode)
			}
		}
	}
	return shortestPath
}

func Part1(in io.Reader) uint16 {
	points := ParseInput(in)
	field := PointsToField(points, 71, 71)
	return FindShortestPath(field, 1024, 71, 71)
}
