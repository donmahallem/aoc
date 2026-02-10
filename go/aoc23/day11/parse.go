package day11

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type point = aoc_utils.Point[int]

type galaxy struct {
	Position point
}

func parseInput(r io.Reader, emptyOffset int) ([]galaxy, error) {
	scanner := bufio.NewScanner(r)
	// map of galaxy positions in x and y
	galaxyMapY := make(map[int]bool, 64)
	galaxyMapX := make(map[int]bool, 64)
	// list of galaxies
	galaxies := make([]galaxy, 0, 32)
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
				galaxies = append(galaxies, galaxy{Position: p})
			} else if line[x] != '.' {
				return nil, aoc_utils.NewUnexpectedInputError(line[x])
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
	return galaxies, nil
}
