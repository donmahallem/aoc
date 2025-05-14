import unittest
import io
from aoc24.day11.part_1 import Part1

testData = "125 17"


class Test2024Day11Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 55312, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()
