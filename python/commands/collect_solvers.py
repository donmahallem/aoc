from .const import SUPPORTED_DAYS, SUPPORTED_PARTS, SUPPORTED_YEARS, Solver
from .get_part import getPart


def collect_solvers() -> list[Solver]:
    solvers: list[Solver] = []
    for year in SUPPORTED_YEARS:
        for day in SUPPORTED_DAYS:
            for part in SUPPORTED_PARTS:
                part_func = getPart(year, day, part)
                if part_func:
                    solvers.append(
                        Solver(year=year, day=day, part=part, func=part_func)
                    )

    return sorted(solvers)
