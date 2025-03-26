package day12

import (
	"fmt"
	"os"
	"slices"

	"github.com/donmahallem/aoc/utils"
)

type VisitedMap map[[2]int]bool

func SortHorizontal(a [2]int, b [2]int) int {
	if a[0] == b[0] {
		return a[1] - b[1]
	}
	return a[0] - b[0]
}

func SortVertical(a [2]int, b [2]int) int {
	if a[1] == b[1] {
		return a[0] - b[0]
	}
	return a[1] - b[1]
}

func CountStraightEdgesHorizontal(coords [][2]int) int {
	slices.SortFunc(coords, SortHorizontal)
	//fmt.Printf("Block %v\n", coords)
	lines := 0
	checkDirs := [2]int{-1, 1}
	var inline bool = false
	for _, checkDir := range checkDirs {
		lastX := coords[0][1]
		lastY := coords[0][0]
		var neighbour [2]int = [2]int{coords[0][0] + checkDir, coords[0][1]}
		inline = !slices.Contains(coords, neighbour)
		for i := 1; i < len(coords); i++ {
			neighbour[0] = coords[i][0] + checkDir
			neighbour[1] = coords[i][1]
			hasNeighbour := slices.Contains(coords, neighbour)
			if inline {
				if hasNeighbour {
					lines++
					inline = false
				} else if utils.Abs(lastX-coords[i][1]) > 1 || lastY != coords[i][0] {
					lines++
					lastX = coords[i][1]
					lastY = coords[i][0]
					inline = !hasNeighbour
				} else {
					lastX = coords[i][1]
					lastY = coords[i][0]
				}
			} else if !hasNeighbour {
				lastX = coords[i][1]
				lastY = coords[i][0]
				inline = true
			}
		}
		if inline {
			lines++
		}
	}
	return lines
}
func CountStraightEdgesVertical(coords [][2]int) int {
	slices.SortFunc(coords, SortVertical)
	//fmt.Printf("Block %v\n", coords)
	lines := 0
	checkDirs := [2]int{-1, 1}
	var inline bool = false
	var neighbour [2]int
	for _, checkDir := range checkDirs {
		lastX := coords[0][1]
		lastY := coords[0][0]
		neighbour[0] = coords[0][0]
		neighbour[1] = coords[0][1] + checkDir
		inline = !slices.Contains(coords, neighbour)
		for i := 1; i < len(coords); i++ {
			neighbour[0] = coords[i][0]
			neighbour[1] = coords[i][1] + checkDir
			hasNeighbour := slices.Contains(coords, neighbour)
			if inline {
				if hasNeighbour {
					lines++
					inline = false
				} else if utils.Abs(lastY-coords[i][0]) > 1 || lastX != coords[i][1] {
					lines++
					lastX = coords[i][1]
					lastY = coords[i][0]
					inline = !hasNeighbour
				} else {
					lastX = coords[i][1]
					lastY = coords[i][0]
				}
			} else if !hasNeighbour && !inline {
				lastX = coords[i][1]
				lastY = coords[i][0]
				inline = true
			}
		}
		if inline {
			lines++
		}
	}
	return lines
}

func CountStraightEdges(coords [][2]int) int {
	edges := 0
	edges += CountStraightEdgesHorizontal(coords)
	edges += CountStraightEdgesVertical(coords)
	return edges
}

func Part2(in *os.File) {
	data, _ := utils.LoadField(in)
	groups := FindGroups(data)
	count := 0
	for _, group := range groups {
		count += len(group) * CountStraightEdges(group)
	}
	fmt.Printf("%d\n", count)
}
