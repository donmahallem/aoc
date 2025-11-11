package day10

import (
	"io"
)

func walkDepthPart2(data Field, x uint16, y uint16, depth uint8) int {
	if depth == 9 {
		return 1
	}
	result := 0
	nextDepth := depth + 1
	centerPoint := y*data.width + x
	if x > 0 && data.Field[centerPoint-1] == nextDepth {
		result += walkDepthPart2(data, x-1, y, nextDepth)
	}
	if y > 0 && data.Field[centerPoint-data.width] == nextDepth {
		result += walkDepthPart2(data, x, y-1, nextDepth)
	}
	if x < data.width-1 && data.Field[centerPoint+1] == nextDepth {
		result += walkDepthPart2(data, x+1, y, nextDepth)
	}
	if y < data.height-1 && data.Field[centerPoint+data.width] == nextDepth {
		result += walkDepthPart2(data, x, y+1, nextDepth)
	}
	return result
}

func searchAllPart2(field Field) int {
	result := 0
	for i := range len(field.starts) {
		result += walkDepthPart2(field, field.starts[i]%field.width, field.starts[i]/field.width, 0)
	}
	return result
}

func Part2(in io.Reader) int {
	data := loadField(in)
	return searchAllPart2(data)
}
