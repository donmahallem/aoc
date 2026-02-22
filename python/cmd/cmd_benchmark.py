import pathlib

from .collect_solvers import collect_solvers
from .const import CommonArgs
from .testdata import TestData
import time
from typing import Dict, Optional, TypedDict
import json
import io

class BenchmarkArgs(CommonArgs):
    year: Optional[int]
    day: Optional[int]
    part: Optional[int]
    timeout: Optional[float]
class ResultStats(TypedDict):
    iterations: int
    avg_ms: float
    total_time_sec: float


def run_benchmark(args: BenchmarkArgs):
    solvers = collect_solvers()

    if args.year is not None:
        solvers = {y: d for y, d in solvers.items() if y == args.year}
    if args.day is not None:
        for y in list(solvers.keys()):
            solvers[y] = {d: p for d, p in solvers[y].items() if d == args.day}
            if not solvers[y]:
                del solvers[y]
    if args.part is not None:
        for y in solvers:
            for d in list(solvers[y].keys()):
                solvers[y][d] = {p: fn for p, fn in solvers[y][d].items() if p == args.part}
                if not solvers[y][d]:
                    del solvers[y][d]

    data_path = pathlib.Path(__file__).parent.parent.parent / "test" / "data.json"
    if not data_path.exists():
        print("Error: Test data not found.")
        return
    test_data = TestData.load(data_path)
    
    timeout_limit = args.timeout if args.timeout else 1.0

    results: Dict[int, Dict[int, Dict[int, Dict[str, ResultStats]]]] = {}

    for year, days in solvers.items():
        if not test_data.has_year(year):
            continue

        for day, parts in days.items():
            if not test_data.has_day(year, day):
                continue

            cases = test_data[year][day]
            for idx, case in enumerate(cases):
                name = case.get("name", f"case{idx}")
                input_data = case.get("input") or (data_path.parent / case["file"]).read_text()

                for part, solver in parts.items():
                    if f"part{part}" not in case:
                        continue

                    iterations = 0
                    elapsed = 0.0
                    bench_start = time.perf_counter()

                    while elapsed < timeout_limit:
                        stream = io.StringIO(input_data)
                        solver(stream)
                        iterations += 1
                        elapsed = time.perf_counter() - bench_start

                    avg_ms = (elapsed / iterations) * 1000 if iterations else 0.0
                    # insert stat into nested results
                    year_dict = results.setdefault(year, {})
                    day_dict = year_dict.setdefault(day, {})
                    part_dict = day_dict.setdefault(part, {})
                    part_dict[name] = {
                        "iterations": iterations,
                        "avg_ms": avg_ms,
                        "total_time_sec": elapsed,
                    }

    if getattr(args, "json", False):
        print(json.dumps(results, indent=2))
    else:
        total_avg = 0.0
        for days in results.values():
            for parts in days.values():
                for cases in parts.values():
                    for stats in cases.values():
                        total_avg += stats['avg_ms']

        print(f"\n{'YEAR':<6} | {'DAY':<4} | {'PART':<4} | {'NAME':<15} | {'ITERATIONS':>12} | {'AVG TIME (ms)':>15} | {'%':>6}")
        print("-" * 103)
        print()  
        for year, days in results.items():
            for day, parts in days.items():
                for part, cases in parts.items():
                    for name, stats in cases.items():
                        pct = (stats['avg_ms'] / total_avg * 100) if total_avg > 0 else 0.0
                        print(f"{year:<6} | {f'{day:02}':<4} | {part:<4} | {name:<15} | {stats['iterations']:>12} | {stats['avg_ms']:>15.4f} | {pct:6.2f}")

