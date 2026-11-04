import argparse

from commands import SolverResult, ListResult
from commands import SUPPORTED_YEARS, SUPPORTED_DAYS, SUPPORTED_PARTS
from commands import CliOutput

if __name__ == "__main__":
    # 1. Root Parser - Global flags go here
    parser = argparse.ArgumentParser(prog="AOC Solver", description="Solves AoC")
    parser.add_argument(
        "--json",
        action=argparse.BooleanOptionalAction,
        help="Output in JSON format (Global)",
    )
    parser.add_argument(
        "--verbose",
        "-v",
        action=argparse.BooleanOptionalAction,
        help="Enable verbose output (Global)",
    )
    parser.add_argument(
        "-o",
        "--output",
        type=str,
        default=None,
        help="Output file for requested data (JSON)",
    )
    sub_parsers = parser.add_subparsers(dest="command", required=True)

    # 2. Solve Subparser
    solve_parser = sub_parsers.add_parser("solve", help="solves a problem")
    solve_parser.add_argument("year", type=int, choices=SUPPORTED_YEARS, help="Year")
    solve_parser.add_argument("day", type=int, choices=SUPPORTED_DAYS, help="Day")
    solve_parser.add_argument("part", type=int, choices=SUPPORTED_PARTS, help="Part")
    solve_parser.add_argument(
        "-i", "--input", type=str, required=True, help="Input file"
    )
    solve_parser.set_defaults(result_class=SolverResult)

    # 3. List Solvers
    list_parser = sub_parsers.add_parser("list", help="lists available solvers")
    list_parser.set_defaults(result_class=ListResult)

    args = parser.parse_args()

    cfg = CliOutput(json=args.json, verbose=args.verbose, output=args.output)
    if hasattr(args, "result_class"):
        result = args.result_class.execute(cfg, args)
        if result is not None:
            cfg.render(result)
