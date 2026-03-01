# Benchmark Aggregation Work

## Overview

The goal is to run benchmarks for both Go and Python implementations on every push to the `main` branch, aggregate their output data into a unified JSON format, and generate a static HTML website that displays a visual history of the performance, separated by the problem (year, day, part, and group, e.g. full vs sample). This website is then published to the `gh-pages` branch.

## Changes Made

### 1. Unified Benchmark Schema

Both Go and Python were updated to emit benchmark metrics in the same schema as defined in `.github/actions/aggregate_benchmark/json_schema.json`.
Both language actions are now dumping `.json` files respectively inside their folders, which are uploaded as GitHub Action Artifacts.

### 2. Modifying `go/convert_benchmark.py`

The original `convert_benchmark.py` produced an incompatible format with missing fields `day`, `year`, `part`.

- It now parses the Go package path (e.g. `github.com/.../aoc/go/aoc24/day01`) to extract `year` and `day`.
- It parses the benchmark name to understand `part` and `group_key` (e.g., `sample` vs `full_data`).
- It outputs exactly the JSON format expected by `json_schema.json`.

### 3. Adding `python/convert_benchmark.py`

Python already natively dumps JSON through `aoc.py benchmark --json`. However, its output was nested recursively (`year -> day -> part -> group -> metrics`).

- We added a `convert_benchmark.py` script specifically for Python to unwrap these recursive dictionaries into a flat array of measurements just as the Go script does.
- It also correctly formats durations into nanoseconds appending the `ns` suffix for compatibility.

### 4. Updating GitHub Actions Workflow (`tests.yml`)

- Removed the old `github-action-benchmark@v1` step which relied on a different structure entirely.
- Added step to run `aoc.py benchmark --json` parsed by `python convert_benchmark.py` to the Python job.
- Added `actions/upload-artifact@v4` steps to both the Go and Python jobs so their results are uploaded to the pipeline and can be retrieved by an aggregator later.
- Added a new `aggregate_benchmarks` job that runs automatically.
- The `aggregate_benchmarks` job downloads the `gh-pages` branch to retain previously generated historical data.
- It downloads all artifacts (using pattern `benchmark-*`).
- It runs the `aggregate.go` action which merges all measurements together into `aggregate.json`.

### 5. Adding `generate_html.py`

To process the aggregated JSON files:

- This script loads the old `data.json` from GitHub Pages, updates it with the new run metrics (overwrites if hash matches to handle re-runs), and computes a consolidated array representing history.
- It then computes the graph datasets and injects them into a simple HTML file using Chart.js configurations.
- The chart groups the benchmarks by Year/Day/Part/Sample combinations, plotting the commits chronologically across the X-axis for both Go and Python lines in exactly the same graphs for comparison.
- The action handles the actual push to the `gh-pages` dataset.
