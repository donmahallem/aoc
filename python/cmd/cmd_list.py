from dataclasses import dataclass
from typing import TYPE_CHECKING, List

from .cli_output import CliOutput

from .const import CommonArgs

if TYPE_CHECKING:
    from .const import Solver


@dataclass(slots=True)
class ListArgs(CommonArgs):
    pass


@dataclass
class ListResult:
    solvers: List['Solver']

    @staticmethod
    def execute(cfg: CliOutput, args: ListArgs) -> 'ListResult':
        from .collect_solvers import collect_solvers

        solvers = collect_solvers()
        return ListResult(solvers=solvers)

    def render_text(self) -> str:
        lines = []
        current_year = None
        printed_days = set()

        for solver in self.solvers:
            if solver.year != current_year:
                if current_year is not None:
                    lines.append("")
                lines.append(f"Year {solver.year}:")
                current_year = solver.year

            # Print once per day
            day_key = (solver.year, solver.day)
            if day_key not in printed_days:
                parts = sorted([
                    s.part for s in self.solvers
                    if s.year == solver.year and s.day == solver.day
                ])
                line = f"  Day {solver.day:02}:"
                for p in parts:
                    line += f" part{p}=true"
                lines.append(line)
                printed_days.add(day_key)

        return "\n".join(lines)

    def to_json(self) -> dict:
        output: dict[int, dict[int, list[int]]] = {}
        for solver in self.solvers:
            y, d, p = solver.year, solver.day, solver.part
            if y not in output:
                output[y] = {}
            if d not in output[y]:
                output[y][d] = []
            if p not in output[y][d]:
                output[y][d].append(p)
        return output
