package day08

import (
	"io"
)

func Part2(in io.Reader) int {
	antennas, width, height := readSource(in)
	var antennaListLen int
	var k int16 = 1
	echos := make(map[Point]bool, 0)
	for antenna := range antennas {
		antennaList := antennas[antenna]
		antennaListLen = len(antennaList)
		for i := range antennaListLen {
			echos[antennaList[i]] = true
			for j := range antennaListLen {
				if i == j {
					continue
				}
				diffX := antennaList[i].X - antennaList[j].X
				diffY := antennaList[i].Y - antennaList[j].Y
				k = 1
				for {
					newX := antennaList[i].X + (k * diffX)
					newY := antennaList[i].Y + (k * diffY)
					if !OutOfBounds(newX, newY, width, height) {
						echos[Point{Y: newY, X: newX}] = true
					} else {
						break
					}
					k++
				}
			}
		}
	}
	return len(echos)
}
