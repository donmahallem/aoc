import unittest
import io
from aoc24.day07.part_1 import Part1

testData = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"""


class Test2024Day07Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 3749, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()
