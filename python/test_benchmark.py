import json
import os
import io
import pytest
from pathlib import Path
from importlib import import_module

CURRENT_DIR = Path(__file__).resolve().parent
TEST_DIR = CURRENT_DIR.parent / "test"

def _get_test_cases():
    data_json_path = CURRENT_DIR.parent / "test" / "data.json"
    if not data_json_path.exists():
        return []

    with open(data_json_path, "r", encoding="utf-8") as f:
        data = json.load(f)

    cases_list = []
    
    for year_str, days in data.items():
        year = int(year_str)
        for day_str, cases in days.items():
            day = int(day_str)
            for i, case in enumerate(cases):
                name = case.get("name", f"case_{i}")
                input_data = case.get("input")
                file_path = case.get("file")

                for part_num, pkey in [(1, "part1"), (2, "part2")]:
                    if pkey in case:
                        expected_data = case[pkey]
                        if "result" in expected_data:
                            cases_list.append({
                                "id": f"{year}_day{day}_{name}_part{part_num}",
                                "year": year,
                                "day": day,
                                "part": part_num,
                                "expected": expected_data["result"],
                                "input_data": input_data,
                                "file_path": file_path
                            })
    return cases_list

def pytest_generate_tests(metafunc):
    if "aoc_case" in metafunc.fixturenames:
        cases = _get_test_cases()
        metafunc.parametrize("aoc_case", cases, ids=[c["id"] for c in cases])

def test_aoc_benchmark(benchmark, aoc_case):
    module_name = f"aoc{aoc_case['year']}.day{aoc_case['day']:02d}.part_{aoc_case['part']}"
    try:
        mod = import_module(module_name)
    except ImportError:
        pytest.skip(f"Module {module_name} not found")

    func_name = f"Part{aoc_case['part']}"
    func = getattr(mod, func_name, None)
    if not func:
        pytest.skip(f"Function {func_name} not found in {module_name}")

    if aoc_case["input_data"] is not None:
        raw_text = aoc_case["input_data"]
    elif aoc_case["file_path"]:
        abs_path = TEST_DIR.joinpath(aoc_case["file_path"]).resolve()
        if not os.path.exists(abs_path):
            pytest.fail(f"Input file not found: {abs_path}")
        with open(abs_path, "r", encoding="utf-8") as f:
            raw_text = f.read()
    else:
        pytest.fail("No input provided for test case")

    # This wrapper isolates the I/O string instantiation, but the main cost is the solution
    def run_part():
        f_obj = io.StringIO(raw_text)
        return func(f_obj)

    # benchmark measures the time it takes to execute run_part
    result = benchmark(run_part)
    assert result == aoc_case["expected"]
