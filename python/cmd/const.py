from typing import Final, Optional, Callable
from dataclasses import dataclass

SUPPORTED_YEARS: Final[tuple[int, ...]] = (24, 25)

SUPPORTED_DAYS: Final[tuple[int, ...]] = tuple(range(1, 26))

SUPPORTED_PARTS: Final[tuple[int, ...]] = (1, 2)

from typing import TypedDict


class CommonArgs(TypedDict):
    json: bool
    verbose: bool


@dataclass
class Solver:
    year: int
    day: int
    part: int
    func: Optional[
        Callable] = None  # None when only listing, populated when executing

    def __lt__(self, other: 'Solver') -> bool:
        """Enable sorting: by year, then day, then part."""
        return (self.year, self.day, self.part) < (other.year, other.day,
                                                   other.part)

