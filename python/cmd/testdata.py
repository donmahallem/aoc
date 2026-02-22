import pathlib
import json
from typing import Dict, TypedDict, List, Optional, Union, Literal, TypeAlias, Mapping

Language: TypeAlias = Literal["go", "python"]
ResultType: TypeAlias = Literal["int", "int16", "string"]

TestResult: TypeAlias = Union[int, str, List[Union[int, str]]]

class PartExpectation(TypedDict, total=False):
    result: TestResult  
    type: ResultType
    skip_languages: List[Language]

class TestCase(TypedDict, total=False):
    name: str           
    input: Optional[str]
    file: Optional[str]
    skip_languages: List[Language]
    part1: PartExpectation
    part2: PartExpectation

ParsedTestData: TypeAlias = Mapping[int, Mapping[int, List[TestCase]]]

class TestData:
    def __init__(self, data: ParsedTestData):
        self._data = data

    @classmethod
    def load(cls, path: str | pathlib.Path) -> "TestData":
        '''Loads test data from a JSON file and parses keys to integers.'''
        path = pathlib.Path(path)
        if not path.exists():
            raise FileNotFoundError(f"Test data file not found: {path}")
        
        raw: Dict[str, Dict[str, List[TestCase]]] = json.loads(path.read_text())
        
        parsed: Dict[int, Dict[int, List[TestCase]]] = {}
        
        for y_str, days in raw.items():
            year = int(y_str)
            parsed[year] = {int(d_str): cases for d_str, cases in days.items()}
            
        return cls(parsed)
    
    def get_year(self, year: int) -> Mapping[int, List[TestCase]]:
        return self._data.get(year, {})

    def get_day(self, year: int, day: int) -> List[TestCase]:
        return self._data.get(year, {}).get(day, [])
    
    def has_year(self, year: int) -> bool:
        return year in self._data
    
    def has_day(self, year: int, day: int) -> bool:
        return day in self._data.get(year, {})

    def __getitem__(self, key: int) -> Mapping[int, List[TestCase]]:
        return self._data[key]

    def __repr__(self) -> str:
        years = list(self._data.keys())
        return f"<TestData years={years}>"