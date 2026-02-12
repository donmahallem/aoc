package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/donmahallem/aoc/go/solver"
)

func TestRunSolveCmd(t *testing.T) {
	t.Run("Solve Day1 Part1", func(t *testing.T) {
		input := "3 4\n4 3\n2 5\n1 3\n3 9\n3 3"
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		ctx := &ExecutionContext{
			Stdout: stdout,
			Stderr: stderr,
			Stdin:  strings.NewReader(input),
			Args:   []string{"24", "1", "1"},
			Solver: solver.NewSolver(),
		}

		exitCode := runSolveCmd(ctx)
		if exitCode != 0 {
			t.Errorf("Exit code = %v, want 0. Output: %s", exitCode, stdout.String())
		}

		output := stdout.String()
		// Take results from sample
		if !strings.Contains(output, "Result is: 11") {
			t.Errorf("Output expected to contain 'Result is: 11', got: %s", output)
		}
	})

	t.Run("Solve NonExisting", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		ctx := &ExecutionContext{
			Stdout: stdout,
			Stderr: stderr,
			Stdin:  strings.NewReader(""),
			Args:   []string{"24", "99", "1"},
			Solver: solver.NewSolver(),
		}

		exitCode := runSolveCmd(ctx)
		if exitCode != 1 {
			t.Errorf("Exit code = %v, want 1", exitCode)
		}
	})

	t.Run("Solve JSON Day1 Part1", func(t *testing.T) {
		input := "3 4\n4 3\n2 5\n1 3\n3 9\n3 3"
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		ctx := &ExecutionContext{
			Stdout: stdout,
			Stderr: stderr,
			Stdin:  strings.NewReader(input),
			Args:   []string{"-json", "24", "1", "1"},
			Solver: solver.NewSolver(),
		}

		exitCode := runSolveCmd(ctx)
		output := stdout.String()

		if exitCode != 0 {
			t.Errorf("Exit code = %v, want 0. Output: %s", exitCode, output)
		}

		var jo JsonOutput
		if err := json.Unmarshal([]byte(output), &jo); err != nil {
			t.Errorf("Output is not valid JSON: %v", err)
		}
		if jo.Result != "11" {
			t.Errorf("JSON Result = %s, want 11", jo.Result)
		}
	})
}

func TestRunSolveCmd_Timeout(t *testing.T) {
	s := solver.NewSolver()

	// Register a slow function
	s.GetRegistry().Register(99, 1, 1, func(in io.Reader) (any, error) {
		time.Sleep(2 * time.Second)
		return "finished", nil
	})

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	ctx := &ExecutionContext{
		Stdout: stdout,
		Stderr: stderr,
		Stdin:  strings.NewReader(""),
		// Set a short timeout of 100ms
		Args:   []string{"-t", "100ms", "99", "1", "1"},
		Solver: s,
	}

	start := time.Now()
	exitCode := runSolveCmd(ctx)
	duration := time.Since(start)

	// Verify duration was roughly the timeout (not 2 seconds)
	if duration > 1500*time.Millisecond {
		t.Errorf("Test took too long (%v), timeout didn't work", duration)
	}

	// Verify exit code
	if exitCode != 1 {
		t.Errorf("Exit code = %v, want 1", exitCode)
	}

	// Verify error message
	expectedError := "solution timed out after 100ms"
	if !strings.Contains(stderr.String(), expectedError) {
		t.Errorf("Stderr expected to contain '%s', got: %s", expectedError, stderr.String())
	}
}
