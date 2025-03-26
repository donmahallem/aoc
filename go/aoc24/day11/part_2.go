package day11

import (
	"fmt"
	"os"
)

func Part2(in *os.File) {
	data, _ := ParseLine(in)

	fmt.Printf("%d\n", SplitStones(data, 75))
}
