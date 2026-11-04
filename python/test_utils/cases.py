"""Benchmark cases built from test/data.json."""

import importlib
from dataclasses import dataclass
from pathlib import Path
from typing import Callable, Optional

from commands.const import SUPPORTED_DAYS, SUPPORTED_YEARS
from commands.testdata import TestData

DATA_PATH = Path(__file__).resolve().parents[2] / "test" / "data.json"


@dataclass(frozen=True)
class BenchCase:
    year: int
    day: int
    part: int
    name: str
    expected: object
    input_text: Optional[str] = None
    input_file: Optional[Path] = None

    @property
    def id(self) -> str:
        return f"{self.year}_day{self.day:02}_{self.name}_part{self.part}"

    def read_input(self) -> str:
        """Raises FileNotFoundError for a referenced input file that is missing."""
        if self.input_text is not None:
            return self.input_text
        assert self.input_file is not None
        return self.input_file.read_text(encoding="utf-8")


def load_cases(
    years: Optional[list[int]] = None,
    days: Optional[list[int]] = None,
    parts: Optional[list[int]] = None,
) -> list[BenchCase]:
    """Python-enabled cases with an expected result; an empty filter selects everything."""
    if not DATA_PATH.exists():
        return []
    data = TestData.load(DATA_PATH)
    cases: list[BenchCase] = []
    for year in SUPPORTED_YEARS:
        if years and year not in years:
            continue
        for day in SUPPORTED_DAYS:
            if days and day not in days:
                continue
            for name, case in data.get_day(year, day).items():
                if "python" in case.get("skip_languages", []):
                    continue
                input_text = case.get("input")
                file_ref = case.get("file")
                if input_text is None and not file_ref:
                    continue
                input_file = DATA_PATH.parent / file_ref if file_ref else None
                for part, expectation in (
                    (1, case.get("part1")),
                    (2, case.get("part2")),
                ):
                    if parts and part not in parts:
                        continue
                    if (
                        expectation is None
                        or "result" not in expectation
                        or "python" in expectation.get("skip_languages", [])
                    ):
                        continue
                    cases.append(
                        BenchCase(
                            year,
                            day,
                            part,
                            name,
                            expectation["result"],
                            input_text,
                            None if input_text is not None else input_file,
                        )
                    )
    return cases


def load_solver(case: BenchCase) -> Optional[Callable]:
    module_name = f"aoc{case.year}.day{case.day:02}.part_{case.part}"
    try:
        module = importlib.import_module(module_name)
    except ModuleNotFoundError as exc:
        missing = exc.name or ""
        if module_name == missing or module_name.startswith(missing + "."):
            return None
        raise
    return getattr(module, f"Part{case.part}")
