package day10

import (
	"io"
)

func WalkDepthPart2(data *Field, x uint16, y uint16, depth uint8) int {
	if depth == 9 {
		return 1
	}
	result := 0
	nextDepth := depth + 1
	if x > 0 && (*data).Field[y][x-1] == nextDepth {
		result += WalkDepthPart2(data, x-1, y, nextDepth)
	}
	if y > 0 && (*data).Field[y-1][x] == nextDepth {
		result += WalkDepthPart2(data, x, y-1, nextDepth)
	}
	if x < (*data).width-1 && (*data).Field[y][x+1] == nextDepth {
		result += WalkDepthPart2(data, x+1, y, nextDepth)
	}
	if y < (*data).height-1 && (*data).Field[y+1][x] == nextDepth {
		result += WalkDepthPart2(data, x, y+1, nextDepth)
	}
	return result
}

func SearchAll2(field *Field) int {
	result := 0
	for i := range len((*field).starts) {
		result += WalkDepthPart2(field, (*field).starts[i].X, (*field).starts[i].Y, 0)
		//result += len(mapper)
	}
	return result
}

func Part2(in io.Reader) int {
	data, _ := LoadField(in)
	return SearchAll2(&data)
}
