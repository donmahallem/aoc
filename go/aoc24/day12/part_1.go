package day12

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Point = aoc_utils.Point[int]

var dirs [4]Point = [4]Point{{Y: 0, X: 1}, {Y: 0, X: -1}, {Y: 1, X: 0}, {Y: -1, X: 0}}

func CountEdges(cells []Point) int {
	edgeCount := 0
	var searchIdx Point = Point{}
	for _, coord := range cells {
		for _, dir := range dirs {
			searchIdx.Y = coord.Y + dir.Y
			searchIdx.X = coord.X + dir.X
			if slices.Contains(cells, searchIdx) {
				continue
			}

			edgeCount += 1
		}
	}
	return edgeCount
}

func findNeighboursInt(field *aoc_utils.ByteField, x int, y int, t byte, found *map[Point]bool) {
	var currentCoord Point = Point{}
	for _, dir := range dirs {
		currentCoord.Y = dir.Y + y
		currentCoord.X = dir.X + x
		if currentCoord.Y < 0 ||
			currentCoord.X < 0 ||
			currentCoord.Y >= int((*field).Height) ||
			currentCoord.X >= int((*field).Width) {
			continue
		} else if (*field).Field[currentCoord.Y][currentCoord.X] == t {
			if _, ok := (*found)[currentCoord]; ok {
				continue
			} else {
				(*found)[currentCoord] = true
				findNeighboursInt(field, currentCoord.X, currentCoord.Y, t, found)
			}
		}
	}
}

func FindNeighbours(field *aoc_utils.ByteField, y int, x int) []Point {
	group := make(map[Point]bool, 0)
	group[Point{Y: y, X: x}] = true
	findNeighboursInt(field, x, y, (*field).Field[y][x], &group)
	keys := make([]Point, 0, len(group))
	for idx := range group {
		keys = append(keys, idx)
	}
	return keys
}

func FindGroups(field *aoc_utils.ByteField) [][]Point {
	taken := make(map[Point]bool)
	var coord Point
	groups := make([][]Point, 0)
	for x := range field.Width {
		for y := range field.Height {
			coord.Y = int(y)
			coord.X = int(x)
			if taken[coord] {
				continue
			}
			neighbours := FindNeighbours(field, int(y), int(x))
			for _, neighbour := range neighbours {
				taken[neighbour] = true
			}
			groups = append(groups, neighbours)
		}
	}
	return groups
}

func Part1(in io.Reader) int {
	data, _ := aoc_utils.LoadField(in)
	groups := FindGroups(data)
	count := 0
	for _, group := range groups {
		count += len(group) * CountEdges(group)
	}
	return count
}
