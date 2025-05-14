import unittest
import io
from aoc24.day10.part_2 import Part2
from .test_part_1 import testData


class Test2024Day10Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 81, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()
