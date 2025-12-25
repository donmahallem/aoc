package day10

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Field struct {
	width, height uint16
	Field         []byte
	starts        []uint16
}

func NewField(width uint16, height uint16, field []byte, starts []uint16) Field {
	return Field{width: width, height: height, Field: field, starts: starts}
}

func loadField(reader io.Reader) (*Field, error) {
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
			// malformed input; return empty field rather than panic
			return nil, aoc_utils.NewParseError("inconsistent line widths in field", nil)
		}
		for idx := uint16(0); idx < uint16(len(line)); idx++ {
			if line[idx] < '0' || line[idx] > '9' {
				return nil, aoc_utils.NewParseError("invalid character in field", nil)
			}
			value := line[idx] - '0'
			obstacles = append(obstacles, value)
			if value == 0 {
				indexPosition := y*width + idx
				starts = append(starts, indexPosition)
			}
		}
		y++
	}
	return &Field{width: width, height: y, Field: obstacles, starts: starts}, nil
}

func walkDepth(data Field, posIdx uint16, depth uint8, ends map[uint16]struct{}) int {
	if depth == 9 {
		// bounds check
		if int(posIdx) >= 0 && int(posIdx) < len(data.Field) {
			ends[posIdx] = struct{}{}
		}
		return 1
	}
	result := 0
	nextDepth := depth + 1

	// left
	if posIdx%data.width > 0 {
		idx := posIdx - 1
		if int(idx) < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepth(data, idx, nextDepth, ends)
		}
	}
	// up
	if posIdx >= data.width {
		idx := posIdx - data.width
		if int(idx) < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepth(data, idx, nextDepth, ends)
		}
	}
	// right
	if (posIdx % data.width) < data.width-1 {
		idx := posIdx + 1
		if int(idx) < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepth(data, idx, nextDepth, ends)
		}
	}
	// down
	if posIdx < (data.height-1)*data.width {
		idx := posIdx + data.width
		if int(idx) < len(data.Field) && data.Field[idx] == nextDepth {
			result += walkDepth(data, idx, nextDepth, ends)
		}
	}
	return result
}

func searchAll(field Field) int {
	result := 0
	for i := 0; i < len(field.starts); i++ {
		mapper := make(map[uint16]struct{})
		walkDepth(field, field.starts[i], 0, mapper)
		result += len(mapper)
	}
	return result
}

func Part1(in io.Reader) (int, error) {
	data, err := loadField(in)
	if err != nil {
		return 0, err
	}
	return searchAll(*data), nil
}
