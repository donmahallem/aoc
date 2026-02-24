import argparse

from cmd import SolverResult, ListResult, BenchmarkResult
from cmd import SUPPORTED_YEARS, SUPPORTED_DAYS, SUPPORTED_PARTS

from cmd import CliOutput, CliOutputConfig

if __name__ == "__main__":
    # 1. Root Parser - Global flags go here
    parser = argparse.ArgumentParser(prog="AOC Solver",
                                     description="Solves and Benchmarks AoC")
    parser.add_argument("--json",
                        action=argparse.BooleanOptionalAction,
                        help="Output in JSON format (Global)")
    parser.add_argument("--verbose",
                        "-v",
                        action=argparse.BooleanOptionalAction,
                        help="Enable verbose output (Global)")

    sub_parsers = parser.add_subparsers(dest="command", required=True)

    # 2. Solve Subparser
    solve_parser = sub_parsers.add_parser("solve", help="solves a problem")
    solve_parser.add_argument("year",
                              type=int,
                              choices=SUPPORTED_YEARS,
                              help="Year")
    solve_parser.add_argument("day",
                              type=int,
                              choices=SUPPORTED_DAYS,
                              help="Day")
    solve_parser.add_argument("part",
                              type=int,
                              choices=SUPPORTED_PARTS,
                              help="Part")
    solve_parser.add_argument("-i",
                              "--input",
                              type=str,
                              required=True,
                              help="Input file")
    solve_parser.set_defaults(result_class=SolverResult)

    # 3. Benchmark Subparser
    bench_parser = sub_parsers.add_parser("benchmark",
                                          help="benchmarks problems")
    bench_parser.add_argument("--year",
                              "-y",
                              type=int,
                              nargs='+',
                              choices=SUPPORTED_YEARS,
                              help="Filter by year(s), e.g. --year 24 25")
    bench_parser.add_argument("--day",
                              "-d",
                              type=int,
                              nargs='+',
                              choices=SUPPORTED_DAYS,
                              help="Filter by day(s), e.g. --day 1 2 24")
    bench_parser.add_argument("--part",
                              "-p",
                              type=int,
                              nargs='+',
                              choices=SUPPORTED_PARTS,
                              help="Filter by part(s)")
    bench_parser.add_argument(
        '-t',
        '--timeout',
        type=float,
        default=1.0,
        help="Maximum time (seconds) to spend on each solver")
    bench_parser.set_defaults(result_class=BenchmarkResult)

    # 4. List Solvers
    list_parser = sub_parsers.add_parser("list",
                                         help="lists available solvers")
    list_parser.set_defaults(result_class=ListResult)

    args = parser.parse_args()

    cfg_dict = CliOutputConfig(json=args.json, verbose=args.verbose)
    cfg = CliOutput(cfg_dict)
    if hasattr(args, 'result_class'):
        result = args.result_class.execute(cfg, args)
        if result is not None:
            cfg.render(result)
