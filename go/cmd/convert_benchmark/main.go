//go:build !wasip1

// convert_benchmark reads the output of
//
//	go test ./... -bench .
//
// from stdin (either plain-text or JSON-lines via --json) and writes a single
// benchmark document (conforming to the aggregate_benchmark JSON schema) to stdout.
//
// Usage:
//
//	go test ./... -bench .        | go run ./cmd/convert_benchmark
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

// timingRe matches the summary line emitted for each benchmark result,
// with optional -benchmem fields:
//
//	797310   1409 ns/op   128 B/op   3 allocs/op
var timingRe = regexp.MustCompile(`(\d+)\s+(\d+(?:\.\d+)?)\s+ns/op(?:\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op)?`)

// pkgRe matches the "pkg: ..." header in plain-text go test output.
var pkgRe = regexp.MustCompile(`^pkg:\s+(\S+)`)

// benchNameRe extracts the full benchmark name from a plain-text benchmark line,
// e.g. "BenchmarkPart1/sample-24   817126   1526 ns/op" → "BenchmarkPart1/sample-24".
var benchNameRe = regexp.MustCompile(`^(Benchmark\S+)\s+`)

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

func parseTest(test string) (part int, sampleName string) {
	switch {
	case strings.Contains(test, "Part1"):
		part = 1
	case strings.Contains(test, "Part2"):
		part = 2
	}
	_, raw, _ := strings.Cut(test, "/")
	// Strip the trailing GOMAXPROCS suffix appended by the Go test runner
	// e.g. "sample-24" → "sample", "full_data-24" → "full_data"
	if idx := strings.LastIndex(raw, "-"); idx >= 0 {
		raw = raw[:idx]
	}
	return part, raw
}

// appendMeasurement builds one Measurement from a package path, test name, and
// the timingRe sub-matches, then appends it to the slice.
func appendMeasurement(measurements []aggBench.Measurement, pkg, test string, m []string) []aggBench.Measurement {
	iters, _ := strconv.Atoi(m[1])
	timingNs, _ := strconv.ParseFloat(m[2], 64)
	duration := fmt.Sprintf("%dns", int64(timingNs))

	year, day := parsePackage(pkg)
	part, sampleName := parseTest(test)

	// description starts as the sample name; benchmem fields may be appended below.
	description := sampleName

	// Append -benchmem fields to description when present (m[3]=B/op, m[4]=allocs/op).
	if m[3] != "" && m[4] != "" {
		mem := fmt.Sprintf("%s B/op, %s allocs/op", m[3], m[4])
		if description != "" {
			description = description + " | " + mem
		} else {
			description = mem
		}
	}

	return append(measurements, aggBench.Measurement{
		SeriesKey:   "go",
		GroupKey:    fmt.Sprintf("%d/%02d/%d_%s", year, day, part, sampleName),
		Duration:    duration,
		Iterations:  iters,
		Description: description,
	})
}

func main() {
	hash := os.Getenv("GITHUB_SHA")
	if hash == "" {
		hash = "unknown"
	}

	var measurements []aggBench.Measurement

	// plainPkg tracks the current package when parsing plain-text go test output.
	var plainPkg string

	scanner := bufio.NewScanner(os.Stdin)
	// increase the buffer size to handle larger benchmarks
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	for scanner.Scan() {
		raw := scanner.Bytes()

		// ── JSON mode ────────────────────────────────────────────────────────
		var ev testEvent
		if err := json.Unmarshal(raw, &ev); err == nil {
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
			measurements = appendMeasurement(measurements, ev.Package, ev.Test, m)
			continue
		}

		// ── Plain-text mode ───────────────────────────────────────────────────
		line := string(raw)

		// Track the package from "pkg: github.com/..." header lines.
		if pm := pkgRe.FindStringSubmatch(line); pm != nil {
			plainPkg = pm[1]
			continue
		}

		// Parse benchmark result lines, e.g.:
		//   BenchmarkPart1/sample-24   817126   1526 ns/op
		nm := benchNameRe.FindStringSubmatch(line)
		if nm == nil {
			continue
		}
		m := timingRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		measurements = appendMeasurement(measurements, plainPkg, nm[1], m)
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
