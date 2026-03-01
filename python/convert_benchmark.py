import json
import os
import sys
from datetime import datetime, timezone
from typing import TextIO


def parse(input_stream: TextIO) -> None:
    try:
        data = json.load(input_stream)
    except json.JSONDecodeError as exc:
        print(f"error: could not parse input JSON: {exc}", file=sys.stderr)
        sys.exit(1)

    measurements: list[dict] = []

    for year_str, year_data in data.items():
        for day_str, day_data in year_data.items():
            for part_str, part_data in day_data.items():
                for group_key, metrics in part_data.items():
                    year = int(year_str)
                    day = int(day_str)
                    part = int(part_str)

                    iterations: int = metrics.get("iterations", 0)
                    avg_ms: float = metrics.get("avg_ms", 0.0)

                    # Convert avg_ms → whole nanoseconds
                    duration_ns = max(1, int(avg_ms * 1_000_000))

                    entry: dict = {
                        "language": "python",
                        "group_key": group_key,
                        "duration": f"{duration_ns}ns",
                        "iterations": iterations,
                        "day": day,
                        "year": year,
                        "part": part,
                    }
                    measurements.append(entry)

    output = {
        "name": "Python Benchmark",
        "hash": os.environ.get("GITHUB_SHA", "unknown"),
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "measurements": measurements,
    }

    json.dump(output, sys.stdout, indent=2)
    sys.stdout.write("\n")


if __name__ == "__main__":
    parse(sys.stdin)
