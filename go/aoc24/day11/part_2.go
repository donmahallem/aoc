package day11

import (
	"fmt"
	"io"
)

func Part2(in io.Reader) {
	data, _ := ParseLine(in)

	fmt.Printf("%d\n", SplitStones(data, 75))
}
