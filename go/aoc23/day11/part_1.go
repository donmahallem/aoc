package day11

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

type point = aoc_utils.Point[int]

type Galaxy struct {
	Position point
}

func ParseInput(r io.Reader, emptyOffset int) []Galaxy {
	scanner := bufio.NewScanner(r)
	// map of galaxy positions in x and y
	galaxyMapY := make(map[int]bool, 64)
	galaxyMapX := make(map[int]bool, 64)
	// list of galaxies
	galaxies := make([]Galaxy, 0, 32)
	var y, maxX, maxY int = 0, 0, 0
	for scanner.Scan() {
		line := scanner.Bytes()
		for x := range line {
			if line[x] == '#' {
				maxX = max(maxX, x)
				maxY = max(maxY, y)
				galaxyMapX[x] = true
				galaxyMapY[y] = true
				p := point{X: x, Y: y}
				galaxies = append(galaxies, Galaxy{Position: p})
			}
		}
		y++
	}

	/*
	 find empty rows and columns between galaxies
	 and add one per column/row to the positions of galaxies beyond the empty row/column
	*/
	galaxyOffsetMapX := make(map[int]int, maxX)
	galaxyOffsetMapY := make(map[int]int, maxY)
	galaxyOffsetMapX[0] = 0
	galaxyOffsetMapY[0] = 0
	for i := range max(maxX, maxY) {
		if i < maxX {
			if _, ok := galaxyMapX[i]; !ok {
				galaxyOffsetMapX[i+1] = galaxyOffsetMapX[i] + emptyOffset
			} else {
				galaxyOffsetMapX[i+1] = galaxyOffsetMapX[i]
			}
		}
		if i < maxY {
			if _, ok := galaxyMapY[i]; !ok {
				galaxyOffsetMapY[i+1] = galaxyOffsetMapY[i] + emptyOffset
			} else {
				galaxyOffsetMapY[i+1] = galaxyOffsetMapY[i]
			}
		}
	}
	// apply offsets to galaxy positions
	for i := range galaxies {
		g := &galaxies[i]
		g.Position.X += galaxyOffsetMapX[g.Position.X]
		g.Position.Y += galaxyOffsetMapY[g.Position.Y]
	}
	return galaxies
}

/**
 * calculates the combined manhatten distances between all galaxies
 */
func combinedManhattenDistances(galaxies []Galaxy) int {

	totalDistance := 0
	totalGalaxies := len(galaxies)
	for idx1 := 0; idx1 < totalGalaxies-1; idx1++ {
		for idx2 := idx1 + 1; idx2 < totalGalaxies; idx2++ {
			g1 := galaxies[idx1]
			g2 := galaxies[idx2]
			totalDistance += int(g1.Position.DistanceManhatten(g2.Position))
		}
	}
	return totalDistance
}

func Part1(in io.Reader) int {
	start := ParseInput(in, 1)
	return combinedManhattenDistances(start)
}
