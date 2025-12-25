package day12

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type point = aoc_utils.Point[int16]
type field = aoc_utils.ByteField[int16, byte]

var dirs [4]point = [4]point{{Y: 0, X: 1}, {Y: 0, X: -1}, {Y: 1, X: 0}, {Y: -1, X: 0}}

func countEdges(cells []point) int {
	edgeCount := 0
	var searchIdx point = point{}
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

func findNeighboursInt(field field, x int16, y int16, t byte, found *map[point]bool) {
	var currentCoord point = point{}
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

func findNeighbours(field field, y int16, x int16) []point {
	group := make(map[point]bool, 8)
	group[point{Y: y, X: x}] = true
	findNeighboursInt(field, x, y, field.Get(x, y), &group)
	keys := make([]point, 0, len(group))
	for idx := range group {
		keys = append(keys, idx)
	}
	return keys
}

func findGroups(field field) [][]point {
	taken := make(map[point]bool, 32)
	var coord point
	groups := make([][]point, 0, 16)
	for x := range field.Width {
		for y := range field.Height {
			coord.Y = y
			coord.X = x
			if taken[coord] {
				continue
			}
			neighbours := findNeighbours(field, y, x)
			for _, neighbour := range neighbours {
				taken[neighbour] = true
			}
			groups = append(groups, neighbours)
		}
	}
	return groups
}

func Part1(in io.Reader) (int, error) {
	data, err := aoc_utils.LoadField[int16, byte](in)
	if err != nil {
		return 0, err
	}
	groups := findGroups(*data)
	count := 0
	for _, group := range groups {
		count += len(group) * countEdges(group)
	}
	return count, nil
}
