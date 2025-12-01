import unittest
import io
from aoc24.day01.part_1 import Part1

testData = """L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"""


class Test2024Day01Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 3, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()
