//go:build !wasip1

// convert_benchmark reads the JSON-lines output produced by
//
//	go test ./... --json -bench .
//
// from stdin and writes a single benchmark document (conforming to the
// aggregate_benchmark JSON schema) to stdout.
//
// NOTE: the --json flag is REQUIRED; plain-text go test output is not supported.
//
// Usage:
//
//	go test ./... --json -bench . | go run ./cmd/convert_benchmark
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	aggBench "github.com/donmahallem/aggregate_benchmark/lib/benchagg"
)

// testEvent is one line of `go test --json` output.
type testEvent struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
}

// timingRe matches the summary line emitted for each benchmark result:
//
//	797310   1409 ns/op
var timingRe = regexp.MustCompile(`(\d+)\s+(\d+(?:\.\d+)?)\s+ns/op`)

func parsePackage(pkg string) (year, day int) {
	parts := strings.Split(pkg, "/")
	if len(parts) < 2 {
		return 0, 0
	}
	yrStr := parts[len(parts)-2]
	dyStr := parts[len(parts)-1]
	if strings.HasPrefix(yrStr, "aoc") {
		year, _ = strconv.Atoi(yrStr[3:])
	}
	if strings.HasPrefix(dyStr, "day") {
		day, _ = strconv.Atoi(dyStr[3:])
	}
	return year, day
}

func parseTest(test string) (part int, groupKey string) {
	switch {
	case strings.Contains(test, "Part1"):
		part = 1
	case strings.Contains(test, "Part2"):
		part = 2
	}
	_, groupKey, _ = strings.Cut(test, "/")
	return part, groupKey
}

func main() {
	hash := os.Getenv("GITHUB_SHA")
	if hash == "" {
		hash = "unknown"
	}

	var measurements []aggBench.Measurement

	scanner := bufio.NewScanner(os.Stdin)
	// increase the buffer size to handle larger benchmarks
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	for scanner.Scan() {
		var ev testEvent
		if err := json.Unmarshal(scanner.Bytes(), &ev); err != nil {
			continue
		}
		if ev.Action != "output" {
			continue
		}
		if ev.Test == "" || !strings.HasPrefix(ev.Test, "Benchmark") {
			continue
		}

		m := timingRe.FindStringSubmatch(ev.Output)
		if m == nil {
			continue
		}
		iters, _ := strconv.Atoi(m[1])
		// timing is in ns; format as integer ns string
		timingNs, _ := strconv.ParseFloat(m[2], 64)
		duration := fmt.Sprintf("%dns", int64(timingNs))

		year, day := parsePackage(ev.Package)
		part, groupKey := parseTest(ev.Test)
		measurements = append(measurements, aggBench.Measurement{
			Language:   "go",
			GroupKey:   groupKey,
			Duration:   duration,
			Iterations: iters,
			Day:        day,
			Year:       year,
			Part:       part,
		})
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	doc := aggBench.BenchmarkFile{
		Name:         "Go Benchmark",
		Hash:         hash,
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
		Measurements: measurements,
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(doc); err != nil {
		fmt.Fprintf(os.Stderr, "encode error: %v\n", err)
		os.Exit(1)
	}
}
