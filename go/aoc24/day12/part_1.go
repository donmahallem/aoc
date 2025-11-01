package day12

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Point = aoc_utils.Point[int16]
type Field = aoc_utils.ByteField[int16, byte]

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

func findNeighboursInt(field Field, x int16, y int16, t byte, found *map[Point]bool) {
	var currentCoord Point = Point{}
	for _, dir := range dirs {
		currentCoord.Y = dir.Y + y
		currentCoord.X = dir.X + x
		if currentCoord.Y < 0 ||
			currentCoord.X < 0 ||
			currentCoord.Y >= field.Height ||
			currentCoord.X >= field.Width {
			continue
		} else if field.Get(currentCoord.X, currentCoord.Y) == t {
			if _, ok := (*found)[currentCoord]; ok {
				continue
			} else {
				(*found)[currentCoord] = true
				findNeighboursInt(field, currentCoord.X, currentCoord.Y, t, found)
			}
		}
	}
}

func FindNeighbours(field Field, y int16, x int16) []Point {
	group := make(map[Point]bool, 8)
	group[Point{Y: y, X: x}] = true
	findNeighboursInt(field, x, y, field.Get(x, y), &group)
	keys := make([]Point, 0, len(group))
	for idx := range group {
		keys = append(keys, idx)
	}
	return keys
}

func FindGroups(field Field) [][]Point {
	taken := make(map[Point]bool, 32)
	var coord Point
	groups := make([][]Point, 0, 16)
	for x := range field.Width {
		for y := range field.Height {
			coord.Y = y
			coord.X = x
			if taken[coord] {
				continue
			}
			neighbours := FindNeighbours(field, y, x)
			for _, neighbour := range neighbours {
				taken[neighbour] = true
			}
			groups = append(groups, neighbours)
		}
	}
	return groups
}

func Part1(in io.Reader) int {
	data, _ := aoc_utils.LoadField[int16, byte](in)
	groups := FindGroups(*data)
	count := 0
	for _, group := range groups {
		count += len(group) * CountEdges(group)
	}
	return count
}
