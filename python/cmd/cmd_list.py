
from .const import CommonArgs
from .collect_solvers import collect_solvers
import json
from .const import SUPPORTED_PARTS


class ListArgs(CommonArgs):
    pass

def run_list(args: ListArgs):
    '''Lists all available solvers.'''
    
    solvers = collect_solvers()
    if args.json:
        out = {y: {d: {p: True for p in parts}
                   for d, parts in days.items()}
               for y, days in solvers.items()}
        print(json.dumps(out))
    else:
        for year, days in solvers.items():
            print(f"Year {year}:")
            for day, parts in days.items():
                line = f"  Day {day:02}:"
                for p in SUPPORTED_PARTS:
                    line += f" part{p}={str(p in parts).lower()}"
                print(line)
