import unittest
import json
import os
import io
from pathlib import Path
from importlib import import_module

CURRENT_DIR = Path(__file__).resolve().parent
TEST_DIR = CURRENT_DIR.parent / "test"


class AOCTestCase(unittest.TestCase):
    """Base class for dynamically generated AOC tests."""

    pass


def _generate_tests():
    """Populate AOCTestCase with test methods based on data.json."""

    # Locate data.json
    data_json_path = CURRENT_DIR.parent / "test" / "data.json"
    if not data_json_path.exists():
        data_json_path = Path("test/data.json").resolve()

    if not data_json_path.exists():
        print(f"Error: Could not find data.json at {data_json_path}")
        return

    try:
        with open(data_json_path, "r", encoding="utf-8") as f:
            data = json.load(f)
    except Exception as e:
        print(f"Error loading data.json: {e}")
        return

    # Create test factory
    def make_test(y, d, p, exp, inp, fp):

        def test_func(self):
            module_name = f"aoc{y}.day{d}.part_{p}"
            try:
                mod = import_module(module_name)
            except ImportError:
                self.skipTest(f"Module {module_name} not found")

            func_name = f"Part{p}"
            func = getattr(mod, func_name, None)
            if not func:
                self.skipTest(f"Function {func_name} not found in {module_name}")

            if inp is not None:
                f_obj = io.StringIO(inp)
            elif fp:
                abs_path = TEST_DIR.joinpath(fp).resolve()
                if not os.path.exists(abs_path):
                    self.fail(f"Input file not found: {abs_path}")
                f_obj = open(abs_path, "r", encoding="utf-8")
            else:
                self.fail("No input provided for test case")

            try:
                result = func(f_obj)
                self.assertEqual(result, exp)
            finally:
                f_obj.close()

        return test_func

    for year, days in data.items():
        for day, cases in days.items():
            for i, case in enumerate(cases):
                name = case.get("name", f"case_{i}")
                input_data = case.get("input")
                file_path = case.get("file")

                parts = []
                if "part1" in case:
                    parts.append((1, case["part1"]))
                if "part2" in case:
                    parts.append((2, case["part2"]))

                if not parts:
                    continue

                for part_num, expected_data in parts:
                    if "result" not in expected_data:
                        continue

                    expected = expected_data["result"]
                    test_name = f"test_{year}_day{day}_{name}_part{part_num}"

                    setattr(
                        AOCTestCase,
                        test_name,
                        make_test(
                            y=year,
                            d=day,
                            p=part_num,
                            exp=expected,
                            inp=input_data,
                            fp=file_path,
                        ),
                    )


_generate_tests()


def load_tests(loader, standard_tests, pattern):
    """unittest discovery protocol - returns a TestSuite."""
    suite = unittest.TestSuite()
    suite.addTests(loader.loadTestsFromTestCase(AOCTestCase))
    return suite


if __name__ == "__main__":
    unittest.main()
