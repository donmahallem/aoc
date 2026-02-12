package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/solver"
)

func TestRunCheckCmd(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantExit int
		wantOut  string
		jsonMode bool
	}{
		{
			name:     "Check existing part",
			args:     []string{"24", "1", "1"},
			wantExit: 0,
			wantOut:  "exists",
		},
		{
			name:     "Check non-existing part",
			args:     []string{"24", "99", "1"},
			wantExit: 1,
			wantOut:  "does not exist",
		},
		{
			name:     "Check JSON existing",
			args:     []string{"-json", "24", "1", "1"},
			wantExit: 0,
			jsonMode: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := solver.NewSolver()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			ctx := &ExecutionContext{
				Stdout: stdout,
				Stderr: stderr,
				Args:   tt.args,
				Solver: s,
			}

			cmd := NewCheckCmd()
			exitCode := cmd.Run(ctx)

			if exitCode != tt.wantExit {
				t.Errorf("runCheckCmd() exit code = %v, want %v", exitCode, tt.wantExit)
			}

			output := stdout.String()
			if tt.jsonMode {
				var jo JsonOutput
				if err := json.Unmarshal([]byte(output), &jo); err != nil {
					t.Errorf("Output is not valid JSON: %v. Output: %s", err, output)
				}
				if jo.Available == nil {
					t.Error("JSON output missing 'available' field")
				}
			} else {
				if !strings.Contains(output, tt.wantOut) {
					t.Errorf("Output = %q, want substring %q", output, tt.wantOut)
				}
			}
		})
	}
}
