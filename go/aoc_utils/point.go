package aoc_utils

import "math"

type PointType interface {
	int | int8 | float32 | float64 | uint16 | int16 | uint | uint8 | int32 | uint32 | int64 | uint64
}
type Point[A PointType] struct {
	X A
	Y A
}

func (a *Point[A]) Diff(b Point[A]) *Point[A] {
	return NewPoint(b.X-a.X, b.Y-a.Y)
}

func (a *Point[A]) DistanceManhatten(b Point[A]) A {
	var dst A = 0
	if b.X < a.X {
		dst += a.X - b.X
	} else {
		dst += b.X - a.X
	}
	if b.Y < a.Y {
		dst += a.Y - b.Y
	} else {
		dst += b.Y - a.Y
	}
	return dst
}

func (a *Point[A]) Add(b Point[A]) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Point[A]) AddXY(x, y A) {
	a.X += x
	a.Y += y
}

func (a *Point[A]) Sub(b Point[A]) {
	a.X -= b.X
	a.Y -= b.Y
}

func (a *Point[A]) SubXY(x, y A) {
	a.X -= x
	a.Y -= y
}

func (a *Point[A]) Multiply(b Point[A]) {
	a.X *= b.X
	a.Y *= b.Y
}

func (a *Point[A]) MultiplyXY(x, y A) {
	a.X *= x
	a.Y *= y
}

// Clone creates a new Point with the same coordinates.
func (a *Point[A]) Clone() *Point[A] {
	return &Point[A]{X: a.X, Y: a.Y}
}

func (a *Point[A]) DistanceEuclid(b Point[A]) float64 {
	return math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2))
}

func NewPoint[A PointType](x, y A) *Point[A] {
	return &Point[A]{x, y}
}

func InBounds[A PointType](x, y, minX, minY, maxX, maxY A) bool {
	return minX > x || minY > y || maxX < x || maxY < y
}

func PointInBounds[A PointType](p *Point[A], minX, minY, maxX, maxY A) bool {
	return InBounds((*p).X, (*p).Y, minX, minY, maxX, maxY)
}
