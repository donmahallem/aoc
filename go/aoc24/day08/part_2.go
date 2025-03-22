package day08

import (
	"fmt"
	"os"
)

func Part2(in *os.File) {
	antennas, width, height := readSource(in)
	var antennaListLen int
	var k int16 = 1
	echos := make(map[[2]int16]bool, 0)
	for antenna := range antennas {
		antennaList := antennas[antenna]
		antennaListLen = len(antennaList)
		for i := range antennaListLen {
			echos[[2]int16{antennaList[i][0], antennaList[i][1]}] = true
			for j := range antennaListLen {
				if i == j {
					continue
				}
				diffX := antennaList[i][1] - antennaList[j][1]
				diffY := antennaList[i][0] - antennaList[j][0]
				k = 1
				for {
					newX := antennaList[i][1] + (k * diffX)
					newY := antennaList[i][0] + (k * diffY)
					if !OutOfBounds(newX, newY, width, height) {
						echos[[2]int16{newY, newX}] = true
					} else {
						break
					}
					k++
				}
			}
		}
	}
	fmt.Printf("%d\n", len(echos))
}
