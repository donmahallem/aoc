package day12

import (
	"fmt"
	"io"
	"slices"

	"github.com/donmahallem/aoc/aoc_utils"
)

var dirs [4][2]int = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func CountEdges(cells [][2]int) int {
	edgeCount := 0
	var searchIdx [2]int = [2]int{0, 0}
	for _, coord := range cells {
		for _, dir := range dirs {
			searchIdx[0] = coord[0] + dir[0]
			searchIdx[1] = coord[1] + dir[1]
			if slices.Contains(cells, searchIdx) {
				continue
			}

			edgeCount += 1
		}
	}
	return edgeCount
}

func findNeighboursInt(field *aoc_utils.ByteField, x int, y int, t byte, found *map[[2]int]bool) {
	var currentCoord [2]int
	for _, dir := range dirs {
		currentCoord[0] = dir[0] + y
		currentCoord[1] = dir[1] + x
		if currentCoord[0] < 0 ||
			currentCoord[1] < 0 ||
			currentCoord[0] >= int((*field).Height) ||
			currentCoord[1] >= int((*field).Width) {
			continue
		} else if (*field).Field[currentCoord[0]][currentCoord[1]] == t {
			if _, ok := (*found)[currentCoord]; ok {
				continue
			} else {
				(*found)[currentCoord] = true
				findNeighboursInt(field, currentCoord[1], currentCoord[0], t, found)
			}
		}
	}
}

func FindNeighbours(field *aoc_utils.ByteField, y int, x int) [][2]int {
	group := make(map[[2]int]bool, 0)
	group[[2]int{y, x}] = true
	findNeighboursInt(field, x, y, (*field).Field[y][x], &group)
	keys := make([][2]int, 0, len(group))
	for idx := range group {
		keys = append(keys, idx)
	}
	return keys
}

func FindGroups(field *aoc_utils.ByteField) [][][2]int {
	taken := make(map[[2]int]bool)
	var coord [2]int
	groups := make([][][2]int, 0)
	for x := range field.Width {
		for y := range field.Height {
			coord[0] = int(y)
			coord[1] = int(x)
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

func Part1(in io.Reader) {
	data, _ := aoc_utils.LoadField(in)
	groups := FindGroups(data)
	count := 0
	for _, group := range groups {
		count += len(group) * CountEdges(group)
	}
	fmt.Printf("%d\n", count)
}
