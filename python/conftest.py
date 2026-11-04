# required for AoC benchmark selection options with pytest
def pytest_addoption(parser):
    group = parser.getgroup("aoc", "AoC benchmark selection")
    for option in ("year", "day", "part"):
        group.addoption(
            f"--aoc-{option}",
            action="append",
            type=int,
            default=[],
            dest=f"aoc_{option}",
            help=f"Only run cases for this {option}; repeat to select several.",
        )
