package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/donmahallem/aoc/go/solver"
)

type CliCommand interface {
	Name() string
	Description() string
	Run(ctx *ExecutionContext) int
	// Usage returns a function that prints the usage information for the command.
	Usage()
}

func parsePartSelector(args []string) (int, int, int, error) {
	if len(args) != 3 {
		return 0, 0, 0, fmt.Errorf("expected <year> <day> <part>")
	}
	year, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid year: %v", err)
	}
	if year >= 2000 {
		year = year % 100
	}
	day, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid day: %v", err)
	}
	part, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid part: %v", err)
	}
	return year, day, part, nil
}

type ExecutionContext struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
	Args   []string
	Solver *solver.Solver
}

type JsonOutput struct {
	Year       int    `json:"year,omitempty"`
	Day        int    `json:"day,omitempty"`
	Part       int    `json:"part,omitempty"`
	Result     string `json:"result,omitempty"`
	DurationUs int64  `json:"duration_us,omitempty"`
	Error      string `json:"error,omitempty"`
	Available  *bool  `json:"available,omitempty"`
}
