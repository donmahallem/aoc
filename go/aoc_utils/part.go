package aoc_utils

import "io"

// PartSelector identifies a specific part of an AoC problem.
type PartSelector struct {
	Year int
	Day  int
	Part int
}

// PartFunc is the generic function signature for a puzzle part.
// It accepts an input reader and returns a result and an error.
type PartFunc[T any] func(in io.Reader) (T, error)

// Adapt converts a strongly typed part function to the generic PartFunc[any].
// This enables type safety without reflection.
func Adapt[T any](fn func(io.Reader) (T, error)) PartFunc[any] {
	return func(in io.Reader) (any, error) {
		return fn(in)
	}
}
