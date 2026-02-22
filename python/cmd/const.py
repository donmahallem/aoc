from typing import Final, TypeVar, Generic, Sequence

SUPPORTED_YEARS: Final[tuple[int, ...]] = (24, 25)

SUPPORTED_DAYS: Final[tuple[int, ...]] = tuple(range(1, 26))

SUPPORTED_PARTS: Final[tuple[int, ...]] = (1, 2)

from typing import TypedDict


class CommonArgs(TypedDict):
    json: bool
