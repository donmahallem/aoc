package day08

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	inputField, err := readSource(in)
	if err != nil {
		return 0, err
	}
	width := int(inputField.Width)
	height := int(inputField.Height)
	// bitmap over the grid; avoids map hashing in the pair loop
	marked := make([]bool, width*height)
	count := 0
	for _, antennaList := range inputField.Antennas {
		for i := range antennaList {
			ax, ay := int(antennaList[i].X), int(antennaList[i].Y)
			if idx := ay*width + ax; !marked[idx] {
				marked[idx] = true
				count++
			}
			for j := range antennaList {
				if i == j {
					continue
				}
				diffX := ax - int(antennaList[j].X)
				diffY := ay - int(antennaList[j].Y)
				for k := 1; ; k++ {
					x := ax + k*diffX
					y := ay + k*diffY
					if x < 0 || y < 0 || x >= width || y >= height {
						break
					}
					if idx := y*width + x; !marked[idx] {
						marked[idx] = true
						count++
					}
				}
			}
		}
	}
	return count, nil
}
