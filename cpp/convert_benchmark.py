"""
Convert Google Benchmark JSON output to the shared aggregate_benchmark format.

Usage:
    aoc_solver_bench.exe --benchmark_format=json | python convert_benchmark.py
    # or from a file:
    python convert_benchmark.py < results.json
    python convert_benchmark.py results.json
"""

import json
import os
import re
import sys
from datetime import datetime, timezone
from typing import TextIO


# Matches:
#   "AoC24/Day01/Part1/full_data/min_time:1.000/repeats:5"
# with optional _mean/_stddev/_cv/_median suffix, and optional /min_time:... /repeats:... segments.
_NAME_RE = re.compile(
    r"^AoC(\d+)/Day(\d+)/Part(\d+)/([^/]+?)(?:/[^_][^/]*)*(?:_mean|_stddev|_median|_cv)?$"
)


def parse(input_stream: TextIO) -> None:
    try:
        data = json.load(input_stream)
    except json.JSONDecodeError as exc:
        print(f"error: could not parse input JSON: {exc}", file=sys.stderr)
        sys.exit(1)

    benchmarks: list[dict] = data.get("benchmarks", [])

    iter_totals: dict[str, int] = {}   # logical_name -> sum of iterations
    mean_rows: dict[str, dict] = {}    # logical_name -> mean aggregate entry
    single_rows: dict[str, dict] = {}  # logical_name -> single-run entry (fallback)

    for b in benchmarks:
        run_type = b.get("run_type", "iteration")
        name: str = b.get("name", "")

        if run_type == "iteration":
            logical = b.get("run_name", name)
            iter_totals[logical] = iter_totals.get(logical, 0) + b.get("iterations", 0)
            if logical not in single_rows:
                single_rows[logical] = b
        elif run_type == "aggregate":
            if b.get("aggregate_name") != "mean":
                continue
            logical = b.get("run_name", re.sub(r"_mean$", "", name))
            mean_rows[logical] = b

    resolved: dict[str, dict] = {}
    for logical in mean_rows:
        resolved[logical] = mean_rows[logical]
    for logical in single_rows:
        if logical not in resolved:
            resolved[logical] = single_rows[logical]

    measurements: list[dict] = []

    for logical_name, b in resolved.items():
        m = _NAME_RE.match(logical_name)
        if not m:
            print(f"warning: could not parse benchmark name: {logical_name!r}", file=sys.stderr)
            continue

        year = int(m.group(1))
        day = int(m.group(2))
        part = int(m.group(3))
        description = m.group(4)

        # cpu_time is already in nanoseconds (time_unit == "ns")
        cpu_time_ns: float = b.get("cpu_time", 0.0)
        duration_ns = max(1, int(cpu_time_ns))

        iterations: int = iter_totals.get(logical_name, b.get("iterations", 0))

        measurements.append(
            {
                "series_key": "cpp",
                "group_key": f"{year}/{day:02d}/{part}_{description}",
                "duration": f"{duration_ns}ns",
                "iterations": iterations,
                "description": description,
            }
        )

    output = {
        "name": "C++ Benchmark",
        "hash": os.environ.get("GITHUB_SHA", "unknown"),
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "measurements": measurements,
    }

    json.dump(output, sys.stdout, indent=2)
    sys.stdout.write("\n")


if __name__ == "__main__":
    if len(sys.argv) > 1:
        with open(sys.argv[1]) as f:
            parse(f)
    else:
        parse(sys.stdin)
