"""pytest-benchmark suite; test_utils/convert_benchmark.py converts its JSON to the aggregate format."""

import io

import pytest

from test_utils.cases import BenchCase, load_cases, load_solver


def pytest_generate_tests(metafunc):
    if "case" not in metafunc.fixturenames:
        return
    option = metafunc.config.getoption
    cases = load_cases(option("aoc_year"), option("aoc_day"), option("aoc_part"))
    metafunc.parametrize("case", cases, ids=[c.id for c in cases])


def test_aoc_benchmark(benchmark, case: BenchCase):
    func = load_solver(case)
    if func is None:
        pytest.skip(f"No solver for {case.year} day {case.day} part {case.part}")
    text = case.read_input()

    benchmark.group = f"{case.year}"
    benchmark.extra_info.update(
        year=case.year, day=case.day, part=case.part, name=case.name
    )

    result = benchmark(lambda: func(io.StringIO(text)))
    assert result == case.expected
