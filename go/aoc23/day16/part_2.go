package day16

import (
	"io"
)

func simulateMax(start Field, memory MovementMemory, pos Position, dir Position, max *int) {
	clear(memory)
	Simulate(start, memory, Movement{Pos: pos, Dir: dir})
	count := countEnergized(memory)
	if count > *max {
		*max = count
	}
}

func Part2(in io.Reader) int {
	start := ParseInputPart1(in)
	max := 0
	memory := make(MovementMemory, start.Width*start.Height)
	// check horizontal borders
	for x := range start.Width {
		simulateMax(start, memory, Position{X: x, Y: 0}, dirDown, &max)
		simulateMax(start, memory, Position{X: x, Y: start.Height - 1}, dirUp, &max)
	}
	// check vertical borders
	for y := range start.Height {
		simulateMax(start, memory, Position{X: 0, Y: y}, dirRight, &max)
		simulateMax(start, memory, Position{X: start.Width - 1, Y: y}, dirLeft, &max)
	}
	return max
}
