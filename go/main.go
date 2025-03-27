package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/donmahallem/aoc/aoc23"
	"github.com/donmahallem/aoc/aoc24"
)

func main() {

	fmt.Print("AOC Solver\n")
	var year, yearError = strconv.ParseUint(os.Args[1], 10, 32)
	if yearError != nil {
		fmt.Println(yearError)
		return
	}
	var day, dayError = strconv.ParseUint(os.Args[2], 10, 32)
	if dayError != nil {
		fmt.Println(dayError)
		return
	}
	var part, partError = strconv.ParseUint(os.Args[3], 10, 32)
	if partError != nil {
		fmt.Println(partError)
		return
	}
	fmt.Printf("Requested parsing %d-%d Part: %d\n", year, day, part)
	switch year {
	case 23:
		aoc23.Aoc23(int(day), int(part))
	case 24:
		aoc24.Aoc24(int(day), int(part))
	}
}
