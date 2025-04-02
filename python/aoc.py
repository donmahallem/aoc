import importlib
import importlib.util
import argparse
import sys

def getPart(year, day, part):
    compound = f"aoc{year}.day{day:02}.part_{part}"
    try:
        spam_spec = importlib.util.find_spec(compound)
        if spam_spec:
            mod = importlib.import_module(compound)
            partName = f"Part{part}"
            if hasattr(mod, partName):
                return getattr(mod, partName)
    except BaseException as error:
        print(error)
        return None

if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        prog="AOC Solver", description="Solves Aoc", epilog="Text at the bottom of help"
    )
    parser.add_argument("year", type=int, choices=[24])  # positional argument
    parser.add_argument(
        "day", type=int, choices=list(range(1, 25))
    )  # option that takes a value
    parser.add_argument("part", type=int, choices=[1, 2])
    args = parser.parse_args()
    print("Requested Year:", args.year, "Day:", args.day, "Part:", args.part)

    solver = getPart(args.year, args.day, args.part)
    if solver:
        solver(sys.stdin)
    else:
        print("Could not find requested solver")
