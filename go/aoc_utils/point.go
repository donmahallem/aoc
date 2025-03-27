package aoc_utils

import "math"

type PointType interface {
	int | int8 | float32 | float64
}
type Point[A PointType] struct {
	X, Y A
}

func (a *Point[A]) Diff(b Point[A]) *Point[A] {
	return NewPoint(b.X-a.X, b.Y-a.Y)
}

func (a *Point[A]) DistanceManhatten(b Point[A]) A {
	return b.X - a.X + b.Y - a.Y
}

func (a *Point[A]) DistanceEuclid(b Point[A]) float64 {
	return math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2))
}

func NewPoint[A PointType](x, y A) *Point[A] {
	return &Point[A]{x, y}
}
