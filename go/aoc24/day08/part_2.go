package day08

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	inputField, err := readSource(in)
	if err != nil {
		return 0, err
	}
	var antennaListLen int
	var k int16 = 1
	var newPoint point
	echos := make(map[point]bool, 100)
	for antenna := range inputField.Antennas {
		antennaList := inputField.Antennas[antenna]
		antennaListLen = len(antennaList)
		for i := 0; i < antennaListLen; i++ {
			echos[antennaList[i]] = true
			for j := 0; j < antennaListLen; j++ {
				if i == j {
					continue
				}
				diffX := antennaList[i].X - antennaList[j].X
				diffY := antennaList[i].Y - antennaList[j].Y
				k = 1
				for {
					newPoint.X = antennaList[i].X + (k * diffX)
					newPoint.Y = antennaList[i].Y + (k * diffY)
					if !OutOfBounds(newPoint.X, newPoint.Y, inputField.Width, inputField.Height) {
						echos[newPoint] = true
					} else {
						break
					}
					k++
				}
			}
		}
	}
	return len(echos), nil
}
