package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/solver"
)

func TestRunListCmd(t *testing.T) {
	t.Run("List Text", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		ctx := &ExecutionContext{
			Stdout: stdout,
			Stderr: stderr,
			Args:   []string{},
			Solver: solver.NewSolver(),
		}

		exitCode := runListCmd(ctx)
		if exitCode != 0 {
			t.Errorf("Exit code = %v, want 0", exitCode)
		}
		if !strings.Contains(stdout.String(), "24") {
			t.Errorf("Output expected to verify 24")
		}
	})

	t.Run("List JSON", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		ctx := &ExecutionContext{
			Stdout: stdout,
			Stderr: stderr,
			Args:   []string{"-json"},
			Solver: solver.NewSolver(),
		}

		exitCode := runListCmd(ctx)
		output := stdout.String()

		if exitCode != 0 {
			t.Errorf("Exit code = %v, want 0", exitCode)
		}

		// The list command returns a specific struct structure, not JsonOutput
		var out struct {
			Years map[string]struct {
				Part1 bool `json:"part1"`
				Part2 bool `json:"part2"`
			} `json:"years"`
		}
		if err := json.Unmarshal([]byte(output), &out); err != nil {
			t.Errorf("Output is not valid JSON: %v. Output: %s", err, output)
		}
		if len(out.Years) == 0 {
			t.Error("JSON output has no years")
		}
	})
}
