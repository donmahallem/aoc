package aoc_utils

import (
	"bufio"
	"bytes"
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils/math"
)

type ByteField[A math.IntType, B math.IntType] struct {
	Width, Height A
	field         []B
}

func (a *ByteField[A, B]) Compare(f *ByteField[A, B]) bool {
	return a.Width == f.Width &&
		a.Height == f.Height &&
		a.field != nil &&
		f.field != nil &&
		slices.Equal(a.field, f.field)
}

func NewField[A math.IntType, B math.IntType](width A, height A, field []B) *ByteField[A, B] {
	return &ByteField[A, B]{Width: width, Height: height, field: field}
}

func (bf *ByteField[A, B]) Get(x A, y A) B {
	return bf.field[y*bf.Width+x]
}

func LoadField[A math.IntType, B math.IntType](reader io.Reader) (*ByteField[A, B], error) {
	return LoadFieldWithOffset[A, B](reader, 0)
}
func LoadFieldWithOffset[A math.IntType, B math.IntType](reader io.Reader, offset byte) (*ByteField[A, B], error) {
	obstacles := make([]B, 0)
	s := bufio.NewScanner(reader)
	y := A(0)
	width := A(0)
	for s.Scan() {
		line := bytes.TrimSuffix(s.Bytes(), []byte{'\r'})
		if width == 0 {
			width = A(len(line))
		} else if width != A(len(line)) {
			panic("Line length is uneven")
		}
		for _, character := range line {
			if offset == 0 {
				obstacles = append(obstacles, B(character))
			} else {
				obstacles = append(obstacles, B(character-offset))
			}
		}
		y++
	}
	return NewField(width, y, obstacles), nil
}
