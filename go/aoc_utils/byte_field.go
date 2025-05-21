package aoc_utils

import (
	"bufio"
	"io"
	"reflect"
)

type ByteField[A IntType] struct {
	Width, Height A
	Field         [][]byte
}

func (a *ByteField[A]) Compare(f *ByteField[A]) bool {
	return a.Width == f.Width &&
		a.Height == f.Height &&
		reflect.DeepEqual(a.Field, f.Field)
}

func NewField[A IntType](width A, height A, field [][]byte) *ByteField[A] {
	return &ByteField[A]{Width: width, Height: height, Field: field}
}

func LoadField[A IntType](reader io.Reader) (*ByteField[A], error) {
	return LoadFieldWithOffset[A](reader, 0)
}
func LoadFieldWithOffset[A IntType](reader io.Reader, offset byte) (*ByteField[A], error) {
	obstacles := make([][]byte, 0)
	s := bufio.NewScanner(reader)
	y := A(0)
	width := A(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = A(len(line))
		} else if width != A(len(line)) {
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
