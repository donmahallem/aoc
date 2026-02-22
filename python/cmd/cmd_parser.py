
from .const import CommonArgs
from .get_part import getPart
import json
import sys

class SolveArgs(CommonArgs):
    year: int
    day: int
    part: int
    input: str
def run_solver(args):
    '''Runs the specified solver with the provided input and outputs the result.'''
    solver = getPart(args.year, args.day, args.part)
    if not solver:
        print("Could not find requested solver")
        return

    if args.input:
        with open(args.input, 'r') as f:
            result = solver(f)
    else:
        result = solver(sys.stdin)

    if args.json:
        print(json.dumps({"year": args.year, "day": args.day, "part": args.part, "result": result}))
    else:
        print(f"Result (Year {args.year}, Day {args.day}, Part {args.part}): {result}")