from .const import SUPPORTED_YEARS, SUPPORTED_DAYS, SUPPORTED_PARTS
from typing import Dict, Optional

from .get_part import getPart
SolverInfoDay = Dict[int, Optional[callable]]
SolverInfo = Dict[int, Dict[int, SolverInfoDay]]
def collect_solvers() -> SolverInfo:
    solvers: SolverInfo = {}
    for year in SUPPORTED_YEARS:
        year_dict: SolverInfoDay = {}
        for day in SUPPORTED_DAYS:
            day_dict: SolverInfoDay = {}
            for part in SUPPORTED_PARTS:
                part_func = getPart(year, day, part)
                if part_func:
                    day_dict[part] = part_func
            if day_dict:
                year_dict[day] = day_dict
        if year_dict:
            solvers[year] = year_dict
    return solvers

