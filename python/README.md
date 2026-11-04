# AOC Python solver

[![codecov](https://codecov.io/gh/donmahallem/aoc/graphs/badge.svg?token=meG36M3hel&component=module_python)](https://app.codecov.io/gh/donmahallem/aoc?components%5B0%5D=Python)

**Usage**

```sh
cat data.txt | python aoc.py {shortYear} {day} {part}
```

## Tests

You can run all tests using:

```sh
python -m unittest -discover
```

Or tests for a specific day with:

```sh
python test_aoc.py -k test_24_day06 -v
```

Before commit please format with

```sh
python -m ruff format ./
```

## Benchmarks

Benchmarks use [pytest-benchmark](https://pytest-benchmark.readthedocs.io) (`pip install -e ".[bench]"`)

```sh
pytest test_benchmark.py --benchmark-only --benchmark-max-time=1.0 --benchmark-min-rounds=1
```

Filter with `--aoc-year`, `--aoc-day` and `--aoc-part` (repeat an option to select several), e.g. `--aoc-year 24 --aoc-day 1 --aoc-day 2`. Standard `-k` also works.

To produce the format used by the benchmark aggregation, convert pytest's JSON with the separate converter (it fails if the input holds no benchmarks):

```sh
pytest test_benchmark.py --benchmark-only --benchmark-json=raw_benchmark.json
python -m test_utils.convert_benchmark raw_benchmark.json > python_benchmark.json
```
