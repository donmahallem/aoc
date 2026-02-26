from .cli_output import CliOutput

from .const import CommonArgs
from typing import Optional, List
import pathlib
import time
import io
from dataclasses import dataclass, asdict


@dataclass
class BenchmarkStats:
    iterations: int
    avg_ms: float
    total_time_sec: float


@dataclass
class BenchmarkEntry:
    year: int
    day: int
    part: int
    name: str
    stats: BenchmarkStats

    def __lt__(self, other: "BenchmarkEntry") -> bool:
        return (self.year, self.day, self.part, self.name) < (
            other.year,
            other.day,
            other.part,
            other.name,
        )


@dataclass(slots=True)
class BenchmarkArgs(CommonArgs):
    year: Optional[list[int]]
    day: Optional[list[int]]
    part: Optional[list[int]]
    timeout: Optional[float]


@dataclass
class BenchmarkResult:
    entries: List["BenchmarkEntry"]

    @staticmethod
    def execute(cfg: CliOutput, args: BenchmarkArgs):
        from .collect_solvers import collect_solvers
        from .testdata import TestData

        solvers = collect_solvers()

        if args.year is not None:
            solvers = [s for s in solvers if s.year in args.year]
        if args.day is not None:
            solvers = [s for s in solvers if s.day in args.day]
        if args.part is not None:
            solvers = [s for s in solvers if s.part in args.part]

        data_path = pathlib.Path(__file__).parent.parent.parent / "test" / "data.json"
        if not data_path.exists():
            cfg.print("Error: Test data not found.")
            return None
        test_data = TestData.load(data_path)

        tasks = []
        for solver in solvers:
            if not test_data.has_year(solver.year):
                continue
            if not test_data.has_day(solver.year, solver.day):
                continue

            cases = test_data[solver.year][solver.day]
            for idx, case in enumerate(cases):
                name = case.get("name", f"case{idx}")
                if "input" in case:
                    input_data = case["input"]
                else:
                    file_ref = case.get("file")
                    if not file_ref:
                        continue
                    input_path = data_path.parent / file_ref
                    if not input_path.exists():
                        cfg.print(f"Warning: Input file {input_path} not found.")
                        continue
                    input_data = input_path.read_text()
                if f"part{solver.part}" in case:
                    tasks.append(
                        {
                            "year": solver.year,
                            "day": solver.day,
                            "part": solver.part,
                            "name": name,
                            "solver": solver.func,
                            "input_data": input_data,
                        }
                    )

        timeout_limit = args.timeout if args.timeout else 1.0
        entries = []

        progress_bar = cfg.progress(tasks, desc="Benchmarking")
        for task in progress_bar:
            year, day, part, name = (
                task["year"],
                task["day"],
                task["part"],
                task["name"],
            )
            solver, input_data = task["solver"], task["input_data"]

            iterations = 0
            elapsed = 0.0
            bench_start = time.perf_counter()

            while elapsed < timeout_limit:
                stream = io.StringIO(input_data)
                solver(stream)
                iterations += 1
                elapsed = time.perf_counter() - bench_start

            avg_ms = (elapsed / iterations) * 1000 if iterations else 0.0

            stats = BenchmarkStats(
                iterations=iterations,
                avg_ms=avg_ms,
                total_time_sec=elapsed,
            )
            entries.append(
                BenchmarkEntry(year=year, day=day, part=part, name=name, stats=stats)
            )

        if hasattr(progress_bar, "close"):
            progress_bar.close()

        return BenchmarkResult(entries=entries)

    def render_text(self) -> str:
        lines = []
        total_avg = sum(e.stats.avg_ms for e in self.entries)

        lines.append("")
        lines.append(
            f"{'YEAR':<6} | {'DAY':<4} | {'PART':<4} | {'NAME':<15} | {'ITERATIONS':>12} | {'AVG TIME (ms)':>15} | {'%':>6}"
        )
        lines.append("-" * 103)
        lines.append("")

        for entry in sorted(self.entries):
            pct = (entry.stats.avg_ms / total_avg * 100) if total_avg > 0 else 0.0
            lines.append(
                f"{entry.year:<6} | {entry.day:02d} | {entry.part:<4} | {entry.name:<15} | {entry.stats.iterations:>12} | {entry.stats.avg_ms:>15.4f} | {pct:6.2f}"
            )

        return "\n".join(lines)

    def to_json(self) -> dict:
        """Convert to JSON-serializable nested dict."""
        output: dict[int, dict[int, dict[int, dict[str, dict]]]] = {}
        for entry in self.entries:
            y, d, p = entry.year, entry.day, entry.part
            if y not in output:
                output[y] = {}
            if d not in output[y]:
                output[y][d] = {}
            if p not in output[y][d]:
                output[y][d][p] = {}
            output[y][d][p][entry.name] = asdict(entry.stats)
        return output
