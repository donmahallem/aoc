package day09

import (
	"io"
)

func Part2(in io.Reader) (int, error) {
	data, width, height, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	positionCache := make(map[int]struct{}, width*5)
	toCheck := make([]int, 0, width*5)
	basins := [4]int{}
	for y := range height {
		for x := range width {
			idx := y*width + x
			value := data[idx]
			if x > 0 && data[idx-1] <= value {
				continue
			}
			if x < width-1 && data[idx+1] <= value {
				continue
			}
			if y > 0 && data[idx-width] <= value {
				continue
			}
			if y < height-1 && data[idx+width] <= value {
				continue
			}
			clear(positionCache)
			clear(toCheck)
			toCheck = append(toCheck, idx)
			for len(toCheck) > 0 {
				current := toCheck[len(toCheck)-1]
				toCheck = toCheck[:len(toCheck)-1]
				if _, ok := positionCache[current]; ok {
					continue
				}
				positionCache[current] = struct{}{}
				cx := current % width
				cy := current / width
				if cx > 0 && data[current-1] != '9' {
					toCheck = append(toCheck, current-1)
				}
				if cx < width-1 && data[current+1] != '9' {
					toCheck = append(toCheck, current+1)
				}
				if cy > 0 && data[current-width] != '9' {
					toCheck = append(toCheck, current-width)
				}
				if cy < height-1 && data[current+width] != '9' {
					toCheck = append(toCheck, current+width)
				}
			}
			basins[3] = len(positionCache)
			// Sort basins in ascending order
			for i := 3; i > 0; i-- {
				if basins[i] > basins[i-1] {
					basins[i], basins[i-1] = basins[i-1], basins[i]
				} else {
					break
				}
			}
		}
	}
	count := 1
	for i := range 3 {
		count *= basins[i]
	}
	return count, nil
}
