package aoc_utils

import (
	"bufio"
	"bytes"
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

type ByteField[A int_util.IntType, B int_util.IntType] struct {
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

// Contains checks if the given coordinates are within the field bounds.
func (bf *ByteField[A, B]) Contains(x, y A) bool {
	return x >= 0 && x < bf.Width && y >= 0 && y < bf.Height
}

// ContainsPoint checks if the given point is within the field bounds.
func (bf *ByteField[A, B]) ContainsPoint(p Point[A]) bool {
	return bf.Contains(p.X, p.Y)
}

func NewField[A int_util.IntType, B int_util.IntType](width A, height A, field []B) *ByteField[A, B] {
	return &ByteField[A, B]{Width: width, Height: height, field: field}
}

// Get safely retrieves the value at x, y. Returns the zero value of B if out of bounds.
func (bf *ByteField[A, B]) Get(x A, y A) B {
	// defensive: protect against out-of-bounds access
	if !bf.Contains(x, y) {
		var zero B
		return zero
	}
	idx := int(y)*int(bf.Width) + int(x)
	return bf.field[idx]
}

// GetPoint safely retrieves the value at the given point.
func (bf *ByteField[A, B]) GetPoint(p Point[A]) B {
	return bf.Get(p.X, p.Y)
}

// Set updates the value at x, y if the coordinates are within bounds.
func (bf *ByteField[A, B]) Set(x A, y A, val B) {
	if bf.Contains(x, y) {
		idx := int(y)*int(bf.Width) + int(x)
		bf.field[idx] = val
	}
}

// SetPoint updates the value at the given point if the coordinates are within bounds.
func (bf *ByteField[A, B]) SetPoint(p Point[A], val B) {
	bf.Set(p.X, p.Y, val)
}

func LoadField[A int_util.IntType, B int_util.IntType](reader io.Reader) (*ByteField[A, B], error) {
	return LoadFieldWithOffset[A, B](reader, 0)
}
func LoadFieldWithOffset[A int_util.IntType, B int_util.IntType](reader io.Reader, offset byte) (*ByteField[A, B], error) {
	obstacles := make([]B, 0)
	s := bufio.NewScanner(reader)
	y := A(0)
	width := A(0)
	for s.Scan() {
		line := bytes.TrimSuffix(s.Bytes(), []byte{'\r'})
		if width == 0 {
			width = A(len(line))
		} else if width != A(len(line)) {
			return nil, NewParseError("line length is uneven", nil)
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
	if width == 0 || y == 0 {
		return nil, NewParseError("empty field", nil)
	}
	if len(obstacles) != int(width)*int(y) {
		return nil, NewParseError("inconsistent field size", nil)
	}
	return NewField(width, y, obstacles), nil
}
