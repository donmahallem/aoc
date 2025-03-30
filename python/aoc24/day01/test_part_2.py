import unittest
import io
from aoc24.day01.part_2 import Part2

testData = """3   4
4   3
2   5
1   3
3   9
3   3"""


class Test2024Day01Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 31, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()
