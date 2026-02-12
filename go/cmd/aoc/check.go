package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type CheckCmd struct {
	fs         *flag.FlagSet
	outputJson *bool
}

func (c CheckCmd) Name() string {
	return "check"
}

func (c CheckCmd) Description() string {
	return "Checks if a solution for the given year, day and part exists"
}

func (c CheckCmd) Usage() {
	fmt.Fprintf(c.fs.Output(), "Usage: %s check [flags] <year> <day> <part>\n", os.Args[0])
	c.fs.PrintDefaults()
}

func (c CheckCmd) Run(ctx *ExecutionContext) int {
	c.fs.SetOutput(ctx.Stderr)
	if err := c.fs.Parse(ctx.Args); err != nil {
		return 1
	}

	y, d, p, err := parsePartSelector(c.fs.Args())
	if err != nil {
		if *c.outputJson {
			out, _ := json.Marshal(JsonOutput{Error: err.Error()})
			fmt.Fprintln(ctx.Stdout, string(out))
		} else {
			fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
			c.fs.Usage()
		}
		return 1
	}

	ok := ctx.Solver.HasSolution(y, d, p)

	if *c.outputJson {
		out, _ := json.Marshal(JsonOutput{
			Year:      y,
			Day:       d,
			Part:      p,
			Available: &ok,
		})
		fmt.Fprintln(ctx.Stdout, string(out))
		return 0
	}

	if ok {
		fmt.Fprintf(ctx.Stdout, "Solution for %d-Day%02d-Part%d exists\n", y, d, p)
		return 0
	}
	fmt.Fprintf(ctx.Stdout, "Solution for %d-Day%02d-Part%d does not exist\n", y, d, p)
	return 1
}

func NewCheckCmd() CliCommand {
	fs := flag.NewFlagSet("check", flag.ExitOnError)
	var jsonF bool
	fs.BoolVar(&jsonF, "json", false, "Output in JSON format")
	fs.BoolVar(&jsonF, "j", false, "Output in JSON format (short)")

	cmd := CheckCmd{fs: fs, outputJson: &jsonF}
	fs.Usage = cmd.Usage
	return cmd
}
