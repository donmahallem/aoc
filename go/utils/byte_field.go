package utils

import (
	"bufio"
	"io"
	"reflect"
)

type ByteField struct {
	Width, Height uint16
	Field         [][]byte
}

func (a *ByteField) Compare(f *ByteField) bool {
	return a.Width == f.Width &&
		a.Height == f.Height &&
		reflect.DeepEqual(a.Field, f.Field)
}

func NewField(width uint16, height uint16, field [][]byte) *ByteField {
	return &ByteField{Width: width, Height: height, Field: field}
}

func LoadField(reader io.Reader) (*ByteField, error) {
	return LoadFieldWithOffset(reader, 0)
}
func LoadFieldWithOffset(reader io.Reader, offset byte) (*ByteField, error) {
	obstacles := make([][]byte, 0)
	s := bufio.NewScanner(reader)
	y := uint16(0)
	width := uint16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = uint16(len(line))
		} else if width != uint16(len(line)) {
			panic("Line length is uneven")
		}
		row := make([]byte, len(line))
		if offset == 0 {
			copy(row, line)
		} else {
			for idx, character := range line {
				row[idx] = character - offset
			}
		}
		obstacles = append(obstacles, row)
		y++
	}
	return NewField(width, y, obstacles), nil
}
