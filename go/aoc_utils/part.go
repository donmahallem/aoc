package aoc_utils

import "io"

type AocPartReturn interface {
	int | []int
}
type AocPart[T AocPartReturn] func(in io.Reader) T
