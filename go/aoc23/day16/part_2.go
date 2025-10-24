package day16

import (
	"io"
)

func clearMemory(mem MovementMemory) {
	for i := range mem {
		mem[i] = 0
	}
}

func Part2(in io.Reader) int {
	start := ParseInputPart1(in)
	max := 0
	memory := make(MovementMemory, start.Width*start.Height)
	// check horizontal borders
	for x := range start.Width {
		clearMemory(memory)
		Simulate(start, memory, Movement{Pos: Position{X: x, Y: 0}, Dir: dirDown})
		count := countEnergized(memory)
		if count > max {
			max = count
		}
	}
	for x := range start.Width {
		clearMemory(memory)
		Simulate(start, memory, Movement{Pos: Position{X: x, Y: start.Height - 1}, Dir: dirUp})
		count := countEnergized(memory)
		if count > max {
			max = count
		}
	}
	// check vertical borders
	for y := range start.Height {
		clearMemory(memory)
		Simulate(start, memory, Movement{Pos: Position{X: 0, Y: y}, Dir: dirRight})
		count := countEnergized(memory)
		if count > max {
			max = count
		}
	}
	for y := range start.Height {
		clearMemory(memory)
		Simulate(start, memory, Movement{Pos: Position{X: start.Width - 1, Y: y}, Dir: dirLeft})
		count := countEnergized(memory)
		if count > max {
			max = count
		}
	}
	return max
}
