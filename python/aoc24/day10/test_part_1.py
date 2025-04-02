import unittest
import io
from aoc24.day10.part_1 import Part1

testData = """89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"""


class Test2024Day10Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 36, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()
