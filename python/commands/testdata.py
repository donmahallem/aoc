import pathlib
import json
from typing import Dict, TypedDict, List, Optional, Union, Literal, Mapping

type Language = Literal["go", "python"]
type ResultType = Literal[
    "int",
    "int16",
    "int32",
    "int64",
    "uint",
    "uint16",
    "uint32",
    "uint64",
    "string",
    "[]int",
    "[]int16",
    "[]int32",
    "[]int64",
    "[]uint",
    "[]uint16",
    "[]uint32",
    "[]uint64",
    "[]string",
    "[]array",
]

type TestResult = Union[int, str, List[Union[int, str]]]


class PartExpectation(TypedDict, total=False):
    result: TestResult
    type: ResultType
    skip_languages: List[Language]


class TestCase(TypedDict, total=False):
    input: Optional[str]
    file: Optional[str]
    skip_languages: List[Language]
    part1: PartExpectation
    part2: PartExpectation


type ParsedTestData = Mapping[int, Mapping[int, Mapping[str, TestCase]]]


class TestData:
    def __init__(self, data: ParsedTestData):
        self._data = data

    @classmethod
    def load(cls, path: str | pathlib.Path) -> "TestData":
        """Loads test data from a JSON file and parses keys to integers."""
        path = pathlib.Path(path)
        if not path.exists():
            raise FileNotFoundError(f"Test data file not found: {path}")

        raw: Dict[str, Dict[str, Dict[str, TestCase]]] = json.loads(path.read_text())

        parsed: Dict[int, Dict[int, Dict[str, TestCase]]] = {}

        for y_str, days in raw.items():
            year = int(y_str)
            parsed[year] = {int(d_str): cases for d_str, cases in days.items()}

        return cls(parsed)

    def get_year(self, year: int) -> Mapping[int, Mapping[str, TestCase]]:
        return self._data.get(year, {})

    def get_day(self, year: int, day: int) -> Mapping[str, TestCase]:
        return self._data.get(year, {}).get(day, {})

    def has_year(self, year: int) -> bool:
        return year in self._data

    def has_day(self, year: int, day: int) -> bool:
        return day in self._data.get(year, {})

    def __getitem__(self, key: int) -> Mapping[int, Mapping[str, TestCase]]:
        return self._data[key]

    def __repr__(self) -> str:
        years = list(self._data.keys())
        return f"<TestData years={years}>"
