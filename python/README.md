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

Filter with `-k`, e.g. `-k 24_day01`.

To produce the format used by the benchmark aggregation:

```sh
pytest test_benchmark.py --benchmark-only --benchmark-max-time=1.0 --benchmark-min-rounds=1 --benchmark-json=raw_benchmark.json
python convert_benchmark.py raw_benchmark.json > python_benchmark.json
```
