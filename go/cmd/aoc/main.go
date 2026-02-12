//go:build !wasip1

package main

import (
	"fmt"
	"os"

	"github.com/donmahallem/aoc/go/solver"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc <command> [<args>]")
		fmt.Println("Commands: solve, check, list")
		os.Exit(1)
	}

	ctx := &ExecutionContext{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
		Args:   os.Args[2:],
		Solver: solver.NewSolver(),
	}

	switch os.Args[1] {
	case "solve":
		os.Exit(runSolveCmd(ctx))
	case "check":
		cmd := NewCheckCmd()
		os.Exit(cmd.Run(ctx))
	case "list":
		os.Exit(runListCmd(ctx))
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		fmt.Println("Usage: aoc <command> [<args>]")
		fmt.Println("Commands: solve, check, list")
		os.Exit(1)
	}
}
