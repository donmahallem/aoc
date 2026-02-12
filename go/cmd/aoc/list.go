//go:build !wasip1

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
)

func runListCmd(ctx *ExecutionContext) int {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	fs.SetOutput(ctx.Stderr)
	var jsonF bool
	fs.BoolVar(&jsonF, "json", false, "Output in JSON format")
	fs.BoolVar(&jsonF, "j", false, "Output in JSON format (short)")

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: %s list [flags]\n", os.Args[0])
		fs.PrintDefaults()
	}

	if err := fs.Parse(ctx.Args); err != nil {
		return 1
	}

	parts := ctx.Solver.GetRegistry().RegisteredParts()

	sort.Slice(parts, func(i, j int) bool {
		if parts[i].Year != parts[j].Year {
			return parts[i].Year < parts[j].Year
		}
		if parts[i].Day != parts[j].Day {
			return parts[i].Day < parts[j].Day
		}
		return parts[i].Part < parts[j].Part
	})

	if jsonF {
		type jsonOutputDay struct {
			Part1 bool `json:"part1"`
			Part2 bool `json:"part2"`
		}
		type days = map[int]jsonOutputDay
		type jsonOutput struct {
			Years map[int]days `json:"years"`
		}
		out := jsonOutput{Years: make(map[int]days)}
		for _, p := range parts {
			if _, ok := out.Years[p.Year]; !ok {
				out.Years[p.Year] = make(days)
			}
			switch p.Part {
			case 1:
				d := out.Years[p.Year][p.Day]
				d.Part1 = true
				out.Years[p.Year][p.Day] = d
			case 2:
				d := out.Years[p.Year][p.Day]
				d.Part2 = true
				out.Years[p.Year][p.Day] = d
			}
		}
		bytes, _ := json.Marshal(out)
		fmt.Fprintln(ctx.Stdout, string(bytes))
		return 0
	}

	var currentYear, currentDay int = -1, -1
	for _, p := range parts {
		if p.Year != currentYear {
			currentYear = p.Year
			fmt.Fprintf(ctx.Stdout, "%d\n", currentYear)
			currentDay = -1
		}
		if p.Day != currentDay {
			currentDay = p.Day
			fmt.Fprintf(ctx.Stdout, "  Day %02d\n", currentDay)
		}
		fmt.Fprintf(ctx.Stdout, "    Part %d\n", p.Part)
	}
	return 0
}
