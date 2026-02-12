package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/donmahallem/aoc/go/solver"
)

func runSolveCmd(ctx *ExecutionContext) int {
	fs := flag.NewFlagSet("solve", flag.ExitOnError)
	fs.SetOutput(ctx.Stderr)
	var inputPath string
	var jsonF bool
	var timeout time.Duration
	fs.StringVar(&inputPath, "input", "", "Path to input file")
	fs.StringVar(&inputPath, "i", "", "Path to input file (short)")
	fs.BoolVar(&jsonF, "json", false, "Output in JSON format")
	fs.BoolVar(&jsonF, "j", false, "Output in JSON format (short)")
	fs.DurationVar(&timeout, "timeout", 10*time.Second, "Timeout for the solution")
	fs.DurationVar(&timeout, "t", 10*time.Second, "Timeout for the solution (short)")

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: %s solve [flags] <year> <day> <part>\n", os.Args[0])
		fs.PrintDefaults()
	}

	if err := fs.Parse(ctx.Args); err != nil {
		return 1
	}

	y, d, p, err := parsePartSelector(fs.Args())
	if err != nil {
		if jsonF {
			out, _ := json.Marshal(JsonOutput{Error: err.Error()})
			fmt.Fprintln(ctx.Stdout, string(out))
		} else {
			fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
			fs.Usage()
		}
		return 1
	}

	if !jsonF {
		fmt.Fprintf(ctx.Stdout, "Requested parsing %d-%d Part: %d\n", y, d, p)
	}

	var input io.Reader = ctx.Stdin
	if inputPath != "" {
		f, err := os.Open(inputPath)
		if err != nil {
			if jsonF {
				out, _ := json.Marshal(JsonOutput{Year: y, Day: d, Part: p, Error: err.Error()})
				fmt.Fprintln(ctx.Stdout, string(out))
			} else {
				fmt.Fprintf(ctx.Stderr, "Failed to open input file: %v\n", err)
			}
			return 1
		}
		defer f.Close()
		input = f
	}

	resultChan := make(chan solver.SolveResult, 1)
	go func() {
		resultChan <- ctx.Solver.Solve(y, d, p, input)
	}()

	var res solver.SolveResult
	select {
	case res = <-resultChan:
	case <-time.After(timeout):
		res = solver.SolveResult{
			Error: fmt.Errorf("solution timed out after %v - likely not solvable or inefficient", timeout),
		}
	}

	if jsonF {
		outObj := JsonOutput{
			Year:       y,
			Day:        d,
			Part:       p,
			DurationUs: res.Duration.Microseconds(),
		}
		if res.Error != nil {
			outObj.Error = res.Error.Error()
		} else {
			outObj.Result = res.Result
		}
		bytes, _ := json.Marshal(outObj)
		fmt.Fprintln(ctx.Stdout, string(bytes))
		return 0
	}

	if res.Error != nil {
		fmt.Fprintln(ctx.Stderr, res.Error)
		return 1
	}

	fmt.Fprintf(ctx.Stdout, "Result is: %s\n", res.Result)
	fmt.Fprintf(ctx.Stdout, "Took: %d\n", res.Duration.Microseconds())
	return 0
}
