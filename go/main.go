//go:build !wasip1

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("AOC Solver\n")
	if len(os.Args) < 4 {
		fmt.Println("Usage: aoc <year> <day> <part>")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	part, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Requested parsing %d-%d Part: %d\n", year, day, part)

	res := RunSolver(year, day, part, os.Stdin)

	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	fmt.Printf("Result is: %s\n", res.Result)
	fmt.Printf("Took: %d\n", res.Duration.Microseconds())
}
