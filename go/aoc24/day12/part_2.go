package day12

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/aoc_utils"
)

type VisitedMap map[Point]bool

func SortHorizontal(a Point, b Point) int {
	if a.Y == b.Y {
		return a.X - b.X
	}
	return a.Y - b.Y
}

func SortVertical(a Point, b Point) int {
	if a.X == b.X {
		return a.Y - b.Y
	}
	return a.X - b.X
}

func CountStraightEdgesHorizontal(coords []Point) int {
	slices.SortFunc(coords, SortHorizontal)
	//fmt.Printf("Block %v\n", coords)
	lines := 0
	checkDirs := [2]int{-1, 1}
	var inline bool = false
	for _, checkDir := range checkDirs {
		lastX := coords[0].X
		lastY := coords[0].Y
		var neighbour Point = Point{Y: coords[0].Y + checkDir, X: coords[0].X}
		inline = !slices.Contains(coords, neighbour)
		for i := 1; i < len(coords); i++ {
			neighbour.Y = coords[i].Y + checkDir
			neighbour.X = coords[i].X
			hasNeighbour := slices.Contains(coords, neighbour)
			if inline {
				if hasNeighbour {
					lines++
					inline = false
				} else if aoc_utils.Abs(lastX-coords[i].X) > 1 || lastY != coords[i].Y {
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
func CountStraightEdgesVertical(coords []Point) int {
	slices.SortFunc(coords, SortVertical)
	//fmt.Printf("Block %v\n", coords)
	lines := 0
	checkDirs := [2]int{-1, 1}
	var inline bool = false
	var neighbour Point
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
				} else if aoc_utils.Abs(lastY-coords[i].Y) > 1 || lastX != coords[i].X {
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

func CountStraightEdges(coords []Point) int {
	edges := 0
	edges += CountStraightEdgesHorizontal(coords)
	edges += CountStraightEdgesVertical(coords)
	return edges
}

func Part2(in io.Reader) int {
	data, _ := aoc_utils.LoadField(in)
	groups := FindGroups(data)
	count := 0
	for _, group := range groups {
		count += len(group) * CountStraightEdges(group)
	}
	return count
}
