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
	// left
	if x > 0 {
		idx := int(centerPoint - 1)
		if idx >= 0 && idx < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepthPart2(data, x-1, y, nextDepth)
		}
	}
	// up
	if y > 0 {
		idx := int(centerPoint - data.width)
		if idx >= 0 && idx < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepthPart2(data, x, y-1, nextDepth)
		}
	}
	// right
	if x < data.width-1 {
		idx := int(centerPoint + 1)
		if idx >= 0 && idx < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepthPart2(data, x+1, y, nextDepth)
		}
	}
	// down
	if y < data.height-1 {
		idx := int(centerPoint + data.width)
		if idx >= 0 && idx < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepthPart2(data, x, y+1, nextDepth)
		}
	}
	return result
}

func searchAllPart2(field Field) int {
	result := 0
	for i := 0; i < len(field.starts); i++ {
		result += walkDepthPart2(field, field.starts[i]%field.width, field.starts[i]/field.width, 0)
	}
	return result
}

func Part2(in io.Reader) (int, error) {
	data, err := loadField(in)
	if err != nil {
		return 0, err
	}
	return searchAllPart2(*data), nil
}
