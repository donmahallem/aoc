"""
Convert pytest-benchmark JSON (--benchmark-json) to the shared aggregate_benchmark format.

Usage (from python/):
    pytest test_benchmark.py --benchmark-only --benchmark-json=raw_benchmark.json
    python -m test_utils.convert_benchmark raw_benchmark.json > python_benchmark.json
    # or from stdin:
    python -m test_utils.convert_benchmark < raw_benchmark.json
"""

import argparse
import json
import os
import sys
from datetime import datetime, timezone


def convert(
    raw: dict, *, series_key: str = "python", bench_name: str = "Python Benchmark"
) -> dict:
    """Raises ValueError for a benchmark without the extra_info set by test_benchmark.py."""
    measurements: list[dict] = []
    for bench in raw.get("benchmarks", []):
        info = bench.get("extra_info", {})
        try:
            year, day, part, name = (
                int(info["year"]),
                int(info["day"]),
                int(info["part"]),
                str(info["name"]),
            )
        except (KeyError, ValueError) as exc:
            raise ValueError(
                f"benchmark {bench.get('name')!r} has invalid extra_info: {exc!r}"
            ) from exc

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


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Convert pytest-benchmark JSON to aggregate format."
    )
    parser.add_argument("file", nargs="?", help="Input JSON file (default: stdin)")
    parser.add_argument("--series-key", default="python", help="Series key")
    parser.add_argument("--name", default="Python Benchmark", help="Suite name")
    args = parser.parse_args()

    try:
        if args.file:
            with open(args.file, encoding="utf-8") as f:
                raw = json.load(f)
        else:
            raw = json.load(sys.stdin)
        output = convert(raw, series_key=args.series_key, bench_name=args.name)
    except (OSError, json.JSONDecodeError, ValueError) as exc:
        print(f"error: {exc}", file=sys.stderr)
        sys.exit(1)

    if not output["measurements"]:
        print("error: no benchmarks in input", file=sys.stderr)
        sys.exit(1)

    json.dump(output, sys.stdout, indent=2)
    sys.stdout.write("\n")


if __name__ == "__main__":
    main()
