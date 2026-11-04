"""pytest-benchmark suite; see README for the flags and conversion to the aggregate format."""

import io
from dataclasses import dataclass
from pathlib import Path

import pytest

from commands import testdata
from commands.const import SUPPORTED_DAYS, SUPPORTED_YEARS
from commands.get_part import getPart

DATA_PATH = Path(__file__).resolve().parent.parent / "test" / "data.json"


@dataclass(frozen=True)
class BenchCase:
    year: int
    day: int
    part: int
    name: str
    text: str
    expected: object

    @property
    def id(self) -> str:
        return f"{self.year}_day{self.day:02}_{self.name}_part{self.part}"


def _load_cases() -> list[BenchCase]:
    if not DATA_PATH.exists():
        return []
    data = testdata.TestData.load(DATA_PATH)
    cases: list[BenchCase] = []
    for year in SUPPORTED_YEARS:
        for day in SUPPORTED_DAYS:
            for name, case in data.get_day(year, day).items():
                if "python" in case.get("skip_languages", []):
                    continue
                text = case.get("input")
                if text is None:
                    file_ref = case.get("file")
                    if not file_ref:
                        continue
                    input_path = DATA_PATH.parent / file_ref
                    if not input_path.exists():
                        continue
                    text = input_path.read_text(encoding="utf-8")
                for part, expectation in (
                    (1, case.get("part1")),
                    (2, case.get("part2")),
                ):
                    if (
                        expectation is None
                        or "result" not in expectation
                        or "python" in expectation.get("skip_languages", [])
                    ):
                        continue
                    cases.append(
                        BenchCase(year, day, part, name, text, expectation["result"])
                    )
    return cases


CASES = _load_cases()


@pytest.mark.parametrize("case", CASES, ids=[c.id for c in CASES])
def test_aoc_benchmark(benchmark, case: BenchCase):
    func = getPart(case.year, case.day, case.part)
    if func is None:
        pytest.skip(f"No solver for {case.year} day {case.day} part {case.part}")

    benchmark.group = f"{case.year}"
    # Read by convert_benchmark.py to build the aggregate keys.
    benchmark.extra_info.update(
        year=case.year, day=case.day, part=case.part, name=case.name
    )

    result = benchmark(lambda: func(io.StringIO(case.text)))
    assert result == case.expected
