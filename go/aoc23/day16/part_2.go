package day16

import (
	"io"
)

func simulateMax(start field, memory movementMemory, pos position, dir position, max *int) {
	clear(memory)
	simulate(start, memory, movement{Pos: pos, Dir: dir})
	count := countEnergized(memory)
	if count > *max {
		*max = count
	}
}

func Part2(in io.Reader) (int, error) {
	start, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	max := 0
	memory := make(movementMemory, start.Width*start.Height)
	// check horizontal borders
	for x := range start.Width {
		simulateMax(*start, memory, position{X: x, Y: 0}, dirDown, &max)
		simulateMax(*start, memory, position{X: x, Y: start.Height - 1}, dirUp, &max)
	}
	// check vertical borders
	for y := range start.Height {
		simulateMax(*start, memory, position{X: 0, Y: y}, dirRight, &max)
		simulateMax(*start, memory, position{X: start.Width - 1, Y: y}, dirLeft, &max)
	}
	return max, nil
}
