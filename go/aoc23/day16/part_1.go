package day16

import (
	"io"
)

func Part1(in io.Reader) (int, error) {
	start, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	memory := make(movementMemory, start.Width*start.Height)
	simulate(*start, memory, movement{Pos: position{X: 0, Y: 0}, Dir: dirRight})
	return countEnergized(memory), nil
}
