package day10

import (
	"bufio"
	"io"
)

type Field struct {
	width, height uint16
	Field         []byte
	starts        []uint16
}

func NewField(width uint16, height uint16, field []byte, starts []uint16) Field {
	return Field{width: width, height: height, Field: field, starts: starts}
}

func loadField(reader io.Reader) Field {
	var obstacles []byte
	starts := make([]uint16, 0)
	s := bufio.NewScanner(reader)
	y := uint16(0)
	width := uint16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = uint16(len(line))
			obstacles = make([]byte, 0, width*width)
		} else if width != uint16(len(line)) {
			panic("Line length is uneven")
		}
		for idx := uint16(0); idx < uint16(len(line)); idx++ {
			value := line[idx] - '0'
			obstacles = append(obstacles, line[idx]-'0')
			if value == 0 {
				indexPosition := y*width + idx
				starts = append(starts, indexPosition)
			}
		}
		y++
	}
	return NewField(width, y, obstacles, starts)
}

func walkDepth(data Field, posIdx uint16, depth uint8, ends map[uint16]struct{}) int {
	if depth == 9 {
		ends[posIdx] = struct{}{}
		return 1
	}
	result := 0
	nextDepth := depth + 1

	if posIdx%data.width > 0 && data.Field[posIdx-1] == nextDepth {
		result += walkDepth(data, posIdx-1, nextDepth, ends)
	}
	if posIdx >= data.width && data.Field[posIdx-data.width] == nextDepth {
		result += walkDepth(data, posIdx-data.height, nextDepth, ends)
	}
	if (posIdx%data.width) < data.width-1 && data.Field[posIdx+1] == nextDepth {
		result += walkDepth(data, posIdx+1, nextDepth, ends)
	}
	if posIdx < (data.height-1)*data.width && data.Field[posIdx+data.width] == nextDepth {
		result += walkDepth(data, posIdx+data.width, nextDepth, ends)
	}
	return result
}

func searchAll(field Field) int {
	result := 0
	mapper := make(map[uint16]struct{})
	for i := range len(field.starts) {
		walkDepth(field, field.starts[i], 0, mapper)
		result += len(mapper)
		clear(mapper)
	}
	return result
}

func Part1(in io.Reader) int {
	data := loadField(in)
	return searchAll(data)
}
