package day12

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

type visitedMap map[point]bool

func sortHorizontal(a point, b point) int {
	if tmp := int(a.Y - b.Y); tmp != 0 {
		return int(tmp)
	}
	return int(a.X - b.X)
}

func sortVertical(a point, b point) int {
	if tmp := int(a.X - b.X); tmp != 0 {
		return tmp
	}
	return int(a.Y - b.Y)

}

func countStraightEdgesHorizontal(coords []point) int {
	slices.SortFunc(coords, sortHorizontal)
	lines := 0
	checkDirs := [2]int16{-1, 1}
	var inline bool = false
	for _, checkDir := range checkDirs {
		lastX := coords[0].X
		lastY := coords[0].Y
		var neighbour point = point{Y: coords[0].Y + checkDir, X: coords[0].X}
		inline = !slices.Contains(coords, neighbour)
		for i := 1; i < len(coords); i++ {
			neighbour.Y = coords[i].Y + checkDir
			neighbour.X = coords[i].X
			hasNeighbour := slices.Contains(coords, neighbour)
			if inline {
				if hasNeighbour {
					lines++
					inline = false
				} else if int_util.AbsInt(lastX-coords[i].X) > 1 || lastY != coords[i].Y {
					lines++
					lastX = coords[i].X
					lastY = coords[i].Y
					inline = !hasNeighbour
				} else {
					lastX = coords[i].X
					lastY = coords[i].Y
				}
			} else if !hasNeighbour {
				lastX = coords[i].X
				lastY = coords[i].Y
				inline = true
			}
		}
		if inline {
			lines++
		}
	}
	return lines
}
func countStraightEdgesVertical(coords []point) int {
	slices.SortFunc(coords, sortVertical)
	//fmt.Printf("Block %v\n", coords)
	lines := 0
	checkDirs := [2]int16{-1, 1}
	var inline bool = false
	var neighbour point
	for _, checkDir := range checkDirs {
		lastX := coords[0].X
		lastY := coords[0].Y
		neighbour.Y = coords[0].Y
		neighbour.X = coords[0].X + checkDir
		inline = !slices.Contains(coords, neighbour)
		for i := 1; i < len(coords); i++ {
			neighbour.Y = coords[i].Y
			neighbour.X = coords[i].X + checkDir
			hasNeighbour := slices.Contains(coords, neighbour)
			if inline {
				if hasNeighbour {
					lines++
					inline = false
				} else if int_util.AbsInt(lastY-coords[i].Y) > 1 || lastX != coords[i].X {
					lines++
					lastX = coords[i].X
					lastY = coords[i].Y
					inline = !hasNeighbour
				} else {
					lastX = coords[i].X
					lastY = coords[i].Y
				}
			} else if !hasNeighbour && !inline {
				lastX = coords[i].X
				lastY = coords[i].Y
				inline = true
			}
		}
		if inline {
			lines++
		}
	}
	return lines
}

func countStraightEdges(coords []point) int {
	edges := 0
	edges += countStraightEdgesHorizontal(coords)
	edges += countStraightEdgesVertical(coords)
	return edges
}

func Part2(in io.Reader) (int, error) {
	data, err := aoc_utils.LoadField[int16, byte](in)
	if err != nil {
		return 0, err
	}
	groups := findGroups(*data)
	count := 0
	for _, group := range groups {
		count += len(group) * countStraightEdges(group)
	}
	return count, nil
}
