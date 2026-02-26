from .cli_output import CliOutput

from .const import CommonArgs
from typing import Optional, Any
from dataclasses import dataclass
import sys

@dataclass(slots=True)
class SolveArgs(CommonArgs):
    year: int
    day: int
    part: int
    input: Optional[str]


@dataclass
class SolverResult:
    year: int
    day: int
    part: int
    result: Any

    @staticmethod
    def execute(cfg: CliOutput, args: SolveArgs) -> Optional['SolverResult']:
        from .get_part import getPart

        solver = getPart(args.year, args.day, args.part)
        if not solver:
            cfg.print("Could not find requested solver")
            return None

        if args.input:
            with open(args.input, 'r') as f:
                result = solver(f)
        else:
            result = solver(sys.stdin)

        return SolverResult(year=args.year,
                            day=args.day,
                            part=args.part,
                            result=result)

    def render_text(self) -> str:
        return f"Result (Year {self.year}, Day {self.day}, Part {self.part}): {self.result}"

    def to_json(self) -> dict:
        return {
            "year": self.year,
            "day": self.day,
            "part": self.part,
            "result": self.result,
        }
