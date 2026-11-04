"""
Convert pytest-benchmark JSON (--benchmark-json) to the shared aggregate_benchmark format.

Usage:
    pytest test_benchmark.py --benchmark-only --benchmark-json=raw_benchmark.json
    python convert_benchmark.py raw_benchmark.json > python_benchmark.json
    # or from stdin:
    python convert_benchmark.py < raw_benchmark.json
"""

import argparse
import json
import os
import sys
from datetime import datetime, timezone
from typing import TextIO


def convert(input_stream: TextIO, *, series_key: str, bench_name: str) -> dict:
    try:
        data = json.load(input_stream)
    except json.JSONDecodeError as exc:
        print(f"error: could not parse input JSON: {exc}", file=sys.stderr)
        sys.exit(1)

    measurements: list[dict] = []
    for bench in data.get("benchmarks", []):
        info = bench.get("extra_info", {})
        try:
            year, day, part, name = (
                int(info["year"]),
                int(info["day"]),
                int(info["part"]),
                str(info["name"]),
            )
        except (KeyError, ValueError):
            print(
                f"warning: missing extra_info for {bench.get('name')!r}",
                file=sys.stderr,
            )
            continue

        stats = bench["stats"]
        duration_ns = max(1, int(stats["mean"] * 1_000_000_000))
        iterations = int(stats["rounds"]) * int(stats.get("iterations", 1))

        measurements.append(
            {
                "series_key": series_key,
                "group_key": f"{year}/{day:02d}/{part}_{name}",
                "duration": f"{duration_ns}ns",
                "iterations": iterations,
                "description": name,
            }
        )

    return {
        "name": bench_name,
        "hash": os.environ.get("GITHUB_SHA", "unknown"),
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "measurements": measurements,
    }


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Convert pytest-benchmark JSON to aggregate format."
    )
    parser.add_argument("file", nargs="?", help="Input JSON file (default: stdin)")
    parser.add_argument("--series-key", default="python", help="Series key")
    parser.add_argument("--name", default="Python Benchmark", help="Suite name")
    args = parser.parse_args()

    if args.file:
        with open(args.file, encoding="utf-8") as f:
            output = convert(f, series_key=args.series_key, bench_name=args.name)
    else:
        output = convert(sys.stdin, series_key=args.series_key, bench_name=args.name)

    json.dump(output, sys.stdout, indent=2)
    sys.stdout.write("\n")
