import unittest
import io
from aoc24.day17.part_1 import Part1

testData = """Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0"""


class Test2024Day15Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(
                result, [4, 6, 3, 5, 6, 3, 5, 2, 1, 0], "The sum is wrong."
            )


if __name__ == "__main__":
    unittest.main()
